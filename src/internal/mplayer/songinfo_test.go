package mplayer

import (
    "os"
    "io/ioutil"
    "path"
    "testing"
    "strings"
)

func TestGetSongInfo(t *testing.T) {
    type TestVector []struct {
        ID3SamplePath string
        ExpectedSongInfo SongInfo
    }
    var testVector TestVector = TestVector{
        { "test-data/dharma_for_one.id3",
          SongInfo {
            "test-data/dharma_for_one.id3",
            "Dharma for One",
            "Jethro Tull",
            "This Was",
            "6",
            "1968",
            "...",
            "Progressive Rock",
          },
        },
        { "test-data/venus_in_force.id3",
          SongInfo {
            "test-data/venus_in_force.id3",
            "Venus In Force",
            "The Hellacopters",
            "Peel Session Live At Maida Vale 04-09-2003",
            "6",
            "2003",
            "",
            "Punk Rock",
          },
        },
        { "test-data/the_electric_index_eel.id3",
          SongInfo {
            "test-data/the_electric_index_eel.id3",
            "The Electric Index Eel",
            "Hellacopters",
            "Grande Rock",
            "5",
            "1999",
            "",
            "Rock",
          },
        },
    }
    _, err := GetSongInfo("blau.mp3")
    if err == nil {
        t.Errorf("err == nil.\n")
    }
    _, err = GetSongInfo("songinfo_test.go")
    if err == nil {
        t.Errorf("err == nil.\n")
    }
    if err.Error() != "Invalid ID3 header." {
        t.Errorf("err has not the expected content : '%s'\n", err.Error())
    }
    wfile, err := os.Create("unsupported-id3.dat")
    if err != nil {
        t.Errorf("err != nil : '%s'\n", err.Error())
    } else {
        defer os.Remove("unsupported-id3.dat")
        wfile.Write([]byte { 'I', 'D', '3', 0x4 })
        wfile.Close()
    }
    _, err = GetSongInfo("")
    if err == nil {
        t.Errorf("err == nil.\n")
    }
    if err.Error() != "Empty song file path was passed." {
        t.Errorf("err has not the expected content : '%s'\n", err.Error())
    }
    for _, testData := range testVector {
        songInfo, err := GetSongInfo(testData.ID3SamplePath)
        if err != nil {
            t.Errorf("songInfo is non-null : '%s'.\n", err.Error())
        }
        if songInfo.FilePath != testData.ExpectedSongInfo.FilePath {
            t.Errorf("songInfo.FilePath has not the expected content.\n")
        }
        if songInfo.Title != testData.ExpectedSongInfo.Title {
            t.Errorf("songInfo.Title has not the expected content : '%s'\n", songInfo.Title)
        }
        if songInfo.Artist != testData.ExpectedSongInfo.Artist {
            t.Errorf("songInfo.Title has not the expected content : '%s'\n", songInfo.Artist)
        }
        if songInfo.Album != testData.ExpectedSongInfo.Album {
            t.Errorf("songInfo.Title has not the expected content : '%s'\n", songInfo.Album)
        }
        if songInfo.TrackNumber != testData.ExpectedSongInfo.TrackNumber {
            t.Errorf("songInfo.Title has not the expected content : '%s'\n", songInfo.TrackNumber)
        }
        if songInfo.Year != testData.ExpectedSongInfo.Year {
            t.Errorf("songInfo.Title has not the expected content : '%s'\n", songInfo.Year)
        }
        if len(testData.ExpectedSongInfo.AlbumCover) != 0 && len(songInfo.AlbumCover) == 0 {
            t.Errorf("songInfo.AlbumCover has no data while it should.\n")
        } else if len(testData.ExpectedSongInfo.AlbumCover) == 0 && len(songInfo.AlbumCover) != 0 {
            t.Errorf("songInfo.AlbumCover has data while it should not.\n")
        }
        if songInfo.Genre != testData.ExpectedSongInfo.Genre {
            t.Errorf("songInfo.Title has not the expected content : '%s'\n", songInfo.Genre)
        }
    }
}

func TestScanSongs(t *testing.T) {
    songs, err := ScanSongs("test-data")
    if err != nil {
        t.Errorf("ScanSongs() has returned an error while it should not.\n")
    }
    if len(songs) != 0 {
        t.Errorf("Slice songs has items while it should not.\n")
    }
    entries, err := os.ReadDir("test-data")
    if err != nil {
        t.Errorf("os.ReadDir() has returned an error while it should not.\n")
    }
    for _, f := range entries {
        if strings.HasSuffix(f.Name(), ".id3") {
            destFilePath := path.Join("test-data", strings.Replace(f.Name(), ".id3", ".mp3", -1))
            data, _ := ioutil.ReadFile(path.Join("test-data", f.Name()))
            ioutil.WriteFile(destFilePath, data, 0644)
            defer os.Remove(destFilePath)
        }
    }
    songs, err = ScanSongs("test-data")
    if err != nil {
        t.Errorf("ScanSongs() has returned an error while it should not.\n")
    }
    if len(songs) != 3 {
        t.Errorf("ScanSongs() has not returned three items as it should.\n")
    }
    _, err = ScanSongs("404-songs")
    if err == nil {
        t.Errorf("ScanSongs() has not returned an error as it should.\n")
    }
}
