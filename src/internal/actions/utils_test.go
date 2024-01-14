package actions

import (
    "testing"
)

func TestGetSongFilePathFromSelectionId(t *testing.T) {
    testVector := []struct {
        Id string
        FilePath string
    }{
        { "jethro tull/this was/my sunday feeling:/abc/my-sunday-feeling.m4a", "/abc/my-sunday-feeling.m4a", },
        { "queens of the stone age/rater-r/regular john:/dev/stoner/regular_john.mp3", "/dev/stoner/regular_john.mp3", },
    }
    for _, test := range testVector {
        filePath := GetSongFilePathFromSelectionId(test.Id)
        if filePath != test.FilePath {
            t.Errorf("Returned filepath different from the expected : '%s' != '%s'.\n", filePath, test.FilePath)
        }
    }
}

func TestGetArtistFromSelectionId(t *testing.T) {
    testVector := []struct {
        Id string
        Artist string
    }{
        { "jethro tull/this was/my sunday feeling:/abc/my-sunday-feeling.m4a", "jethro tull", },
        { "queens of the stone age/rater-r/regular john:/dev/stoner/regular_john.mp3", "queens of the stone age", },
    }
    for _, test := range testVector {
        artist := GetArtistFromSelectionId(test.Id)
        if artist != test.Artist {
            t.Errorf("Returned artist different from the expected : '%s' != '%s'.\n", artist, test.Artist)
        }
    }
}

func TestGetAlbumFromSelectionId(t *testing.T) {
    testVector := []struct {
        Id string
        Album string
    }{
        { "jethro tull/this was/my sunday feeling:/abc/my-sunday-feeling.m4a", "this was", },
        { "queens of the stone age/rater-r/regular john:/dev/stoner/regular_john.mp3", "rater-r", },
    }
    for _, test := range testVector {
        album := GetAlbumFromSelectionId(test.Id)
        if album != test.Album {
            t.Errorf("Returned album different from the expected : '%s' != '%s'.\n", album, test.Album)
        }
    }
}