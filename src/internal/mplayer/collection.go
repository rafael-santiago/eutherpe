package mplayer

type MusicCollection map[string]map[string][]SongInfo

func LoadMusicCollection(basePath string) (MusicCollection, error) {
    songs, err := ScanSongs(basePath)
    if err != nil {
        return nil, err
    }
    collection := make(MusicCollection)
    for _, song := range songs {
        album, has := collection[song.Artist]
        if !has {
            album = make(map[string][]SongInfo)
        }
        album[song.Album] = append(album[song.Album], song)
        collection[song.Artist] = album
    }
    return collection, nil
}
