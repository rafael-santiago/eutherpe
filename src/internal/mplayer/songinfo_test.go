package mplayer

import (
    _"os"
    _"path"
    "testing"
)

func TestGetSongInfo(t *testing.T) {
    //wd, _ := os.Getwd()
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
