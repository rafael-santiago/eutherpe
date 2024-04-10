package mplayer

import (
    "testing"
    "os"
    "io/ioutil"
    "strings"
    "path"
)

func TestLoadMusicCollection(t *testing.T) {
    collection, err := LoadMusicCollection("/tmp")
    if err != nil {
        t.Errorf("LoadMusicCollection() has returned an error : '%s'\n", err.Error())
    }
    if len(collection) != 0 {
        t.Errorf("LoadMusicCollection() has not returned an empty collection.\n");
    }
    entries, err := os.ReadDir("test-data")
    if err != nil {
        t.Errorf("os.ReadDir() has returned an error while it should not.\n")
    }
    for _, f := range entries {
        if strings.HasSuffix(f.Name(), ".id3") {
            destFilePath := path.Join("/tmp", strings.Replace(f.Name(), ".id3", ".mp3", -1))
            data, _ := ioutil.ReadFile(path.Join("test-data", f.Name()))
            ioutil.WriteFile(destFilePath, data, 0644)
            defer os.Remove(destFilePath)
        }
    }
    collection, err = LoadMusicCollection("/tmp")
    if len(collection) != 3 {
        t.Errorf("LoadMusicCollection() has returned a wrong total of items.\n")
    }
    _, err = LoadMusicCollection("404-songs")
    if err == nil {
        t.Errorf("LoadMusicCollection() has not returned an error when it should.\n")
    }
}

func TestGetArtistsFromCollection(t *testing.T) {
    collection := make(MusicCollection)
    collection["Queens Of The Stone Age"] = make(map[string][]SongInfo)
    collection["Motorhead"] = make(map[string][]SongInfo)
    collection["The Cramps"] = make(map[string][]SongInfo)
    collection["Queens Of The Stone Age"]["Queens Of The Stone Age"] = []SongInfo {
        SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", },
    }
    collection["Motorhead"]["Bomber"] = []SongInfo {
        SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", },
    }
    collection["Motorhead"]["Overkill"] = []SongInfo {
        SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
    }
    collection["The Cramps"]["Songs The Lord Taught Us"] = []SongInfo {
        SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    artists := GetArtistsFromCollection(collection)
    if len(artists) != 3 {
        t.Errorf("The returned slice does not have three items as expected.\n")
    }
    if artists[0] != "Motorhead" ||
       artists[1] != "Queens Of The Stone Age" ||
       artists[2] != "The Cramps" {
        t.Errorf("The slice seems not to be sorted : '%v'.\n", artists)
    }
}

func TestGetAlbumsFromArtist(t *testing.T) {
    collection := make(MusicCollection)
    collection["Queens Of The Stone Age"] = make(map[string][]SongInfo)
    collection["Motorhead"] = make(map[string][]SongInfo)
    collection["The Cramps"] = make(map[string][]SongInfo)
    collection["Queens Of The Stone Age"]["Queens Of The Stone Age"] = []SongInfo {
        SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", },
    }
    collection["Queens Of The Stone Age"]["Songs For The Deaf"] = []SongInfo {
        SongInfo { "no-one-knows.mp3", "No One Knows", "Queens Of The Stone Age", "Songs For The Deaf", "2", "2002", "", "Stoner Rock", },
    }
    collection["Queens Of The Stone Age"]["In Times New Roman"] = []SongInfo {
        SongInfo { "emotion-sickness.mp3", "Emotion Sickness", "Queens Of The Stone Age", "In Times New Roman", "9", "2023", "", "Stoner Rock" },
    }
    collection["Motorhead"]["Bomber"] = []SongInfo {
        SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", },
    }
    collection["Motorhead"]["Overkill"] = []SongInfo {
        SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
    }
    collection["The Cramps"]["Songs The Lord Taught Us"] = []SongInfo {
        SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    motorhead_albums := GetAlbumsFromArtist("Motorhead", collection)
    if len(motorhead_albums) != 2 {
        t.Errorf("Returned slice of albums does not have the expected length.\n")
    }
    if motorhead_albums[0] != "Overkill" ||
       motorhead_albums[1] != "Bomber" {
        t.Errorf("Returned slice was not sorted as expected : '%v'.\n", motorhead_albums)
    }
    the_cramps_albums := GetAlbumsFromArtist("The Cramps", collection)
    if len(the_cramps_albums) != 1 {
        t.Errorf("Returned slice of albums does not have the expected length.\n")
    }
    if the_cramps_albums[0] != "Songs The Lord Taught Us" {
        t.Errorf("Returned slice of albums does not have the expected content : '%v'.\n", the_cramps_albums)
    }
    qotsa_albums := GetAlbumsFromArtist("Queens Of The Stone Age", collection)
    if len(qotsa_albums) != 3 {
        t.Errorf("Returned slice of albums does not have the expected length.\n")
    }
    if qotsa_albums[0] != "In Times New Roman" ||
       qotsa_albums[1] != "Songs For The Deaf" ||
       qotsa_albums[2] != "Queens Of The Stone Age" {
        t.Errorf("Returned slice of albums was not sorted as expected : '%v'.\n", qotsa_albums)
    }
    artist404 := GetAlbumsFromArtist("artist404", collection)
    if len(artist404) != 0 {
        t.Errorf("Unknown artist has returned albums : '%v'.\n", artist404)
    }
}

func TestGetSongFromArtistAlbum(t *testing.T) {
    collection := make(MusicCollection)
    collection["Queens Of The Stone Age"] = make(map[string][]SongInfo)
    collection["Motorhead"] = make(map[string][]SongInfo)
    collection["The Cramps"] = make(map[string][]SongInfo)
    collection["Queens Of The Stone Age"]["Queens Of The Stone Age"] = []SongInfo {
        SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", },
    }
    collection["Queens Of The Stone Age"]["Songs For The Deaf"] = []SongInfo {
        SongInfo { "no-one-knows.mp3", "No One Knows", "Queens Of The Stone Age", "Songs For The Deaf", "2", "2002", "", "Stoner Rock", },
    }
    collection["Queens Of The Stone Age"]["In Times New Roman"] = []SongInfo {
        SongInfo { "emotion-sickness.mp3", "Emotion Sickness", "Queens Of The Stone Age", "In Times New Roman", "9", "2023", "", "Stoner Rock" },
    }
    collection["Motorhead"]["Bomber"] = []SongInfo {
        SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", },
    }
    collection["Motorhead"]["Overkill"] = []SongInfo {
        SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
    }
    collection["The Cramps"]["Songs The Lord Taught Us"] = []SongInfo {
        SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    _, err := collection.GetSongFromArtistAlbum("Falcao",
                                                "Dinheiro nao eh tudo, mas eh 70%",
                                                "Atirei o Pau No Gato (Soh Para Ouvir o Miau).m4a")
    if err == nil {
        t.Errorf("GetSongFromArtistAlbum() did not return an error.\n")
    }
    if err.Error() != "No collection for Falcao." {
        t.Errorf("Unexpected error message.\n")
    }
    _, err = collection.GetSongFromArtistAlbum("Queens Of The Stone Age", "...Like Clockwork", "i-sat-by-the-ocean.mp3")
    if err == nil {
        t.Errorf("GetSongFromArtistAlbum() did not return an error.\n")
    }
    if err.Error() != "No album ...Like Clockwork for Queens Of The Stone Age was found." {
        t.Errorf("Unexpected error message.\n")
    }
    _, err = collection.GetSongFromArtistAlbum("Queens Of The Stone Age", "Queens Of The Stone Age", "the-bronze.mp3")
    if err == nil {
        t.Errorf("GetSongFromArtistAlbum() did not return an error.\n")
    }
    if err.Error() != "No song the-bronze.mp3 in album Queens Of The Stone Age by Queens Of The Stone Age was found." {
        t.Errorf("Unexpected error message.\n")
    }
    for artist, albums := range collection {
        for album, songs := range albums {
            for _, song := range songs {
                curr_song, err := collection.GetSongFromArtistAlbum(artist, album, song.FilePath)
                if err != nil {
                    t.Errorf("GetSongFromArtistAlbum() returned an error.\n")
                }
                if curr_song != song {
                    t.Errorf("curr_song != song : '%v' != '%v'.\n", curr_song, song)
                }
            }
        }
    }
}
