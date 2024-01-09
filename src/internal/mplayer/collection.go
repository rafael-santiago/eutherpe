package mplayer

import(
    "sort"
    "strings"
    "strconv"
    "fmt"
)

type MusicCollection map[string]map[string][]SongInfo

func (m *MusicCollection) GetSongFromArtistAlbum(artist, album, filePath string) (SongInfo, error) {
    artistCollection, has := (*m)[artist]
    if !has {
        return SongInfo{}, fmt.Errorf("No collection for %s.", artist)
    }
    songs, has := artistCollection[album]
    if !has {
        return SongInfo{}, fmt.Errorf("No album %s for %s was found.", album, artist)
    }
    for _, song := range songs {
        if song.FilePath == filePath {
            return song, nil
        }
    }
    return SongInfo{}, fmt.Errorf("No song %s in album %s by %s was found.", filePath, album, artist)
}

func LoadMusicCollection(basePath string) (MusicCollection, error) {
    songs, err := ScanSongs(basePath)
    if err != nil {
        return nil, err
    }
    collection := make(MusicCollection)
    for _, song := range songs {
        albums, has := collection[song.Artist]
        if !has {
            albums = make(map[string][]SongInfo)
        }
        albums[song.Album] = append(albums[song.Album], song)
        collection[song.Artist] = albums
    }
    for _, albums := range collection {
        for album, tracks := range albums {
            albums[album] = sortTracksFromAlbum(tracks)
        }
    }
    return collection, nil
}

func GetArtistsFromCollection(collection MusicCollection) []string {
    artists := make([]string, 0)
    for artist, _ := range collection {
        if !has(artists, artist) {
            artists = append(artists, artist)
        }
    }
    sort.Strings(artists)
    return artists
}

func GetAlbumsFromArtist(artist string, collection MusicCollection) []string {
    albums, exists := collection[artist]
    if !exists {
        return make([]string, 0)
    }
    albumsByYear := make(map[string]string)
    for album, song := range albums {
        albumsByYear[guessUpAlbumYear(song) + "_" + album] = album
    }
    var years []string
    for year, _ := range albumsByYear {
        years = append(years, year)
    }
    sort.Slice(years, func (i, j int) bool { return years[i] > years[j] })
    var albumsFromPresentToPast []string
    for _, yearAlbum := range years {
        i := strings.Index(yearAlbum, "_")
        albumsFromPresentToPast = append(albumsFromPresentToPast, yearAlbum[i+1:])
    }
    return albumsFromPresentToPast
}

func sortTracksFromAlbum(trackList []SongInfo) []SongInfo {
    trackNumbers := make([]string, 0)
    for _, track := range trackList {
        trackNumbers = append(trackNumbers, track.TrackNumber)
    }
    albumTracks := make([]SongInfo, 0)
    sort.Slice(trackNumbers,
               func (i, j int) bool {
                    a, _ := strconv.Atoi(trackNumbers[i])
                    b, _ := strconv.Atoi(trackNumbers[j])
                    return a < b
                })
    for _, trackNumber := range trackNumbers {
        for _, track := range trackList {
            if track.TrackNumber == trackNumber {
                albumTracks = append(albumTracks, track)
            }
        }
    }
    return albumTracks
}

func guessUpAlbumYear(albumTracks []SongInfo) string {
    years := make(map[string]int)
    for _, track := range albumTracks {
        if len(track.Year) > 0 {
            years[track.Year] += 1
        }
    }
    foundFreq := 0
    foundYear := "Unk"
    for year, freq := range years {
        if freq > foundFreq {
            foundYear = year
            foundFreq = freq
        }
    }
    return foundYear
}

func has(haystack []string, needle string) bool {
    for _, h := range haystack {
        if (h == needle) {
            return true
        }
    }
    return false
}

