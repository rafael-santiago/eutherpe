package mplayer

import(
    "sort"
    "strings"
    "strconv"
    "fmt"
    "encoding/json"
    "encoding/base64"
    "os"
)

type MusicCollection map[string]map[string][]SongInfo

type AlbumsArray struct {
    Album string
    Songs []SongInfo
}

type ArtistsArray struct {
    Artist string
    Albums []AlbumsArray
}

type CollectionArray struct {
    Artists []ArtistsArray
}

type WholeDeviceCollection struct {
    Collection CollectionArray
}

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

func LoadMusicCollection(basePath string, coversCacheRootPath ...string) (MusicCollection, error) {
    songs, err := ScanSongs(basePath, coversCacheRootPath...)
    if err != nil {
        return nil, err
    }
    collection := make(MusicCollection)
    for _, song := range songs {
        albums, has := collection[song.Artist]
        if !has {
            albums = make(map[string][]SongInfo)
        }
        found := false
        for _, previousSong := range albums[song.Album] {
            found = (previousSong.Title == song.Title)
            if found {
                break
            }
        }
        if !found {
            albums[song.Album] = append(albums[song.Album], song)
            collection[song.Artist] = albums
        }
    }
    for _, albums := range collection {
        for album, tracks := range albums {
            albums[album] = sortTracksFromAlbum(tracks)
            albumCover := ""
            for _, song := range tracks {
                if len(albumCover) == 0 && len(song.AlbumCover) > 0 {
                    albumCover = song.AlbumCover
                    break
                }
            }
            if len(albumCover) == 0 {
                continue
            }
            for t, _ := range albums[album] {
                if len(albums[album][t].AlbumCover) == 0 {
                    albums[album][t].AlbumCover = albumCover
                }
            }
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

func (m *MusicCollection) ToJSON() string {
    var deviceCollection WholeDeviceCollection
    artistsFromCollection := GetArtistsFromCollection(*m)
    artistsFromCollectionLen := len(artistsFromCollection)
    if artistsFromCollectionLen == 0 {
        return ""
    }
    deviceCollection.Collection.Artists = make([]ArtistsArray, artistsFromCollectionLen)
    for a, currentArtist := range artistsFromCollection {
        albumsFromArtist := GetAlbumsFromArtist(currentArtist, *m)
        deviceCollection.Collection.Artists[a].Artist = currentArtist
        deviceCollection.Collection.Artists[a].Albums = getAlbumsAsArray(currentArtist, albumsFromArtist, *m)
    }
    data, err := json.Marshal(deviceCollection)
    if err != nil {
        return ""
    }
    return string(data)
}

func (m *MusicCollection) FromJSON(filePath string) error {
    var deviceCollection WholeDeviceCollection
    fileData, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }
    err = json.Unmarshal([]byte(fileData), &deviceCollection)
    if err != nil {
        return err
    }
    (*m) = make(MusicCollection)
    for _, artist := range deviceCollection.Collection.Artists {
        (*m)[artist.Artist] = make(map[string][]SongInfo)
        for _, album := range artist.Albums {
            albumCover := ""
            for _, song := range album.Songs {
                isCachedCover := strings.HasPrefix(song.AlbumCover, "blob-id=")
                if !isCachedCover && len(song.AlbumCover) > 0 {
                    blob, _ := base64.StdEncoding.DecodeString(song.AlbumCover)
                    albumCover = string(blob)
                    break
                } else if isCachedCover {
                    albumCover = song.AlbumCover
                    break
                }
            }
            (*m)[artist.Artist][album.Album] = album.Songs
            for s, _ := range (*m)[artist.Artist][album.Album] {
                (*m)[artist.Artist][album.Album][s].AlbumCover = albumCover
            }
        }
    }
    return nil
}

func sortTracksFromAlbum(trackList []SongInfo) []SongInfo {
    trackNumber := len(trackList) + 1
    for t, _ := range trackList {
        if len(trackList[t].TrackNumber) == 0 {
            trackList[t].TrackNumber = fmt.Sprintf("%s", trackNumber)
            trackNumber++
        }
    }
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

func getAlbumsAsArray(currentArtist string, albumsFromArtist[]string, collection MusicCollection) []AlbumsArray {
    albumsArray := make([]AlbumsArray, len(albumsFromArtist))
    for a, currentAlbum := range albumsFromArtist {
        albumsArray[a].Album = currentAlbum
        albumsArray[a].Songs = getSongsAsArray(currentArtist, currentAlbum, collection)
    }
    return albumsArray
}

func getSongsAsArray(artist string, album string, collection MusicCollection) []SongInfo {
    albumSongs := collection[artist][album]
    songs := make([]SongInfo, len(albumSongs))
    for s, currSong := range albumSongs {
        songs[s] = currSong
        isCachedCover := strings.HasPrefix(currSong.AlbumCover, "blob-id=")
        if !isCachedCover && len(currSong.AlbumCover) > 0 {
            songs[s].AlbumCover = base64.StdEncoding.EncodeToString([]byte(currSong.AlbumCover))
        }
    }
    return songs
}
