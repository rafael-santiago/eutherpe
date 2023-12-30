package dj

import (
    "testing"
    "internal/mplayer"
    "os"
)

func TestAdd(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
        playlist.Add(song)
    }
}

func TestGetSongByFilePath(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
    }
    for _, song := range songs {
        s, err := playlist.GetSongByFilePath(song.FilePath)
        if err != nil {
            t.Errorf("GetSongByFilePath() has returned an error when it must not : '%s'.\n", err.Error())
        } else if s != song {
            t.Errorf("Element returned by GetSongByFilePath has unexpected data.\n")
        }
    }
    _, err := playlist.GetSongByFilePath("FlorentinaEhONomeDela")
    if err == nil {
        t.Errorf("GetSongByFilePath has not returned an error when it must.\n")
    } else if err.Error() != "'FlorentinaEhONomeDela' not found in playlist 'TestPlaylist'." {
        t.Errorf("GetSongByFilePath has returned an unexpected error message.\n")
    }
}

func TestRemove(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
    }
    for _, song := range songs {
        playlist.Remove(song)
        _, err := playlist.GetSongByFilePath(song.FilePath)
        if err == nil {
            t.Errorf("Remove() seems like not being removing what it should.\n")
        }
        playlist.Remove(song)
    }
}

func TestGetSongIndexByFilePath(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
    }
    for s, song := range songs {
        if playlist.GetSongIndexByFilePath(song.FilePath) != s {
            t.Errorf("GetSongIndexByFilePath() has returned a wrong index.\n")
        }
    }
    if playlist.GetSongIndexByFilePath("Ziriguidum!") != -1 {
        t.Errorf("GetSongIndexByFilePath() has not returned -1 when it supposed must.\n")
    }
}

func TestMoveUp(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
    }
    if playlist.GetSongIndexByFilePath(songs[2].FilePath) != 2 {
        t.Errorf("GetSongIndexByFilePath() has returned wrong index.\n")
    }
    playlist.MoveUp(songs[2])
    if playlist.GetSongIndexByFilePath(songs[2].FilePath) != 1 {
        t.Errorf("MoveUp() seems not work.\n")
    }
    if playlist.GetSongIndexByFilePath(songs[1].FilePath) != 2 {
        t.Errorf("MoveUp() seems not work.\n")
    }
    playlist.MoveUp(songs[2])
    if playlist.GetSongIndexByFilePath(songs[2].FilePath) != 0 {
        t.Errorf("MoveUp() seems not work.\n")
    }
    if playlist.GetSongIndexByFilePath(songs[0].FilePath) != 1 {
        t.Errorf("MoveUp() seems not work.\n")
    }
    playlist.MoveUp(songs[2])
    if playlist.GetSongIndexByFilePath(songs[2].FilePath) != 0 {
        t.Errorf("MoveUp() seems not work.\n")
    }
}

func TestMoveDown(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
    }
    if playlist.GetSongIndexByFilePath(songs[0].FilePath) != 0 {
        t.Errorf("GetSongIndexByFilePath() has returned wrong index.\n")
    }
    playlist.MoveDown(songs[0])
    if playlist.GetSongIndexByFilePath(songs[0].FilePath) != 1 {
        t.Errorf("MoveDown() seems not work.\n")
    }
    if playlist.GetSongIndexByFilePath(songs[1].FilePath) != 0 {
        t.Errorf("MoveDown() seems not work.\n")
    }
    playlist.MoveDown(songs[0])
    if playlist.GetSongIndexByFilePath(songs[0].FilePath) != 2 {
        t.Errorf("MoveDown() seems not work.\n")
    }
    if playlist.GetSongIndexByFilePath(songs[2].FilePath) != 1 {
        t.Errorf("MoveDown() seems not work.\n")
    }
    playlist.MoveDown(songs[0])
    if playlist.GetSongIndexByFilePath(songs[0].FilePath) != 2 {
        t.Errorf("MoveDown() seems not work.\n")
    }
}

func TestSaveToLoadFrom(t *testing.T) {
    songs := []mplayer.SongInfo {
        {
            "/mnt/jukebox/feeling-good.mp3",
            "Feeling Good",
            "Nina Simone",
            "Best of Nina Simone",
            "10",
            "1999",
            "some data...",
            "Jazz/Soul",
        },
        {
            "/mnt/jukebox/lawman.mp3",
            "Lawnman",
            "Motorhead",
            "Bomber",
            "2",
            "1979",
            "",
            "Speed Metal",
        },
        {
            "/mnt/jukebox/mambo sun.mp3",
            "Mambo Sun",
            "T-Rex",
            "Electric Warrior",
            "1",
            "1971",
            "(...)",
            "Glam Rock",
        },
    }
    playlist := Playlist{}
    playlist.Name = "TestPlaylist"
    for _, song := range songs {
        playlist.Add(song)
    }
    err := playlist.SaveTo("")
    if err == nil {
        t.Errorf("SaveTo() has succeeded when it must fail.\n")
    }
    err = playlist.SaveTo("TestPlaylist.eu")
    if err != nil {
        t.Errorf("SaveTo has failed.\n")
    }
    defer os.Remove("TestPlaylist.eu")
    err = playlist.LoadFrom("TestPlaylist")
    if err == nil {
        t.Errorf("LoadFrom() has succeeded when it must fail.\n")
    }
    playlistFromDisk := Playlist{}
    err = playlistFromDisk.LoadFrom("TestPlaylist.eu")
    if err != nil {
        t.Errorf("LoadFrom() has failed.\n")
    }
    for s, song := range songs {
        if playlistFromDisk.GetSongIndexByFilePath(song.FilePath) != s {
            t.Errorf("Loaded SongInfo has not the expected index inside the list.\n")
        }
        songFromDisk, err := playlistFromDisk.GetSongByFilePath(song.FilePath)
        if err != nil {
            t.Errorf("GetSongByFilePath() has failed.\n")
        }
        if songFromDisk != song {
            t.Errorf("SongInfo loaded from disk seems corrupted.\n")
        }
    }
}
