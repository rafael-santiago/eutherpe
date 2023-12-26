package mplayer

import (
    "testing"
    "os"
    "io/ioutil"
    "strings"
    "path"
)

func TestLoadMusicCollection(t *testing.T) {
    collection, err := LoadMusicCollection(".")
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
            destFilePath := path.Join("test-data", strings.Replace(f.Name(), ".id3", ".mp3", -1))
            data, _ := ioutil.ReadFile(path.Join("test-data", f.Name()))
            ioutil.WriteFile(destFilePath, data, 0644)
            defer os.Remove(destFilePath)
        }
    }
    collection, err = LoadMusicCollection(".")
    if len(collection) != 3 {
        t.Errorf("LoadMusicCollection() has returned a wrong total of items.\n")
    }
    _, err = LoadMusicCollection("404-songs")
    if err == nil {
        t.Errorf("LoadMusicCollection() has not returned an error when it should.\n")
    }
}
