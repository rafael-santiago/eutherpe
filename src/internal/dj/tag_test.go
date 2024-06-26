//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package dj

import (
    "testing"
    "os"
)

func TestTagAdd(t *testing.T) {
    tags := &Tags{}
    tags.Add("i'm in the band", "rock")
    tags.Add("red house", "RoCk", "blues")
    tags.Add("soul serenade", "soul")
    tags.Add("baby please don't go", "blues")
    tags.Add("feeling good", "jazz")
    if len(tags.Tags) != 4 {
        t.Errorf("Wrong amount of tags.\n")
    }
}

func TestTagGet(t *testing.T) {
    tags := &Tags{}
    tags.Add("i'm in the band", "rock")
    tags.Add("red house", "rock", "Blues")
    tags.Add("soul serenade", "soul")
    tags.Add("baby please don't go", "blues")
    tags.Add("feeling good", "jazz")
    songs := tags.Get("rock")
    if len(songs) != 2 {
        t.Errorf("Wrong amount of rock songs.")
    } else if songs[0] != "i'm in the band" ||
              songs[1] != "red house" {
        t.Errorf("Unexpected rock songs.\n")
    }
    songs = tags.Get("blues")
    if len(songs) != 2 {
        t.Errorf("Wrong amount of blues songs.\n")
    } else if songs[0] != "red house" ||
              songs[1] != "baby please don't go" {
        t.Errorf("Unexpected blues songs.\n")
    }
    songs = tags.Get("soul")
    if len(songs) != 1 {
        t.Errorf("Wrong amount of soul songs.\n")
    } else if songs[0] != "soul serenade" {
        t.Errorf("Unexpected soul song.\n")
    }
    songs = tags.Get("jazz")
    if len(songs) != 1 {
        t.Errorf("Wrong amount of jazz songs.\n")
    } else if songs[0] != "feeling good" {
        t.Errorf("Unexpected jazz song.\n")
    }
    songs = tags.Get("BregaSofrenciaExtremo")
    if len(songs) != 0 {
        t.Errorf("Unexistent tag has returned songs!!!\n")
    }
}

func TestTagDel(t *testing.T) {
    tags := &Tags{}
    tags.Add("i'm in the band", "rOck")
    tags.Add("red house", "rock", "blues")
    tags.Add("baby please don't go", "bLues")
    songs := tags.Get("rock")
    if len(songs) != 2 {
        t.Errorf("Wrong amount of rock songs.\n")
    } else if songs[0] != "i'm in the band" ||
              songs[1] != "red house" {
        t.Errorf("Unexpected rock songs.\n")
    }
    songs = tags.Get("blues")
    if len(songs) != 2 {
        t.Errorf("Wrong amount of blues songs.\n")
    } else if songs[0] != "red house" ||
              songs[1] != "baby please don't go" {
        t.Errorf("Unexpected blues songs.\n")
    }
    tags.Del("red house", "blues", "rock")
    songs = tags.Get("rock")
    if len(songs) != 1 {
        t.Errorf("Wrong amount of rock songs.\n")
    } else if songs[0] != "i'm in the band" {
        t.Errorf("Unexpected rock song.\n")
    }
    songs = tags.Get("blues")
    if len(songs) != 1 {
        t.Errorf("Wrong amount of blues songs.\n")
    } else if songs[0] != "baby please don't go" {
        t.Errorf("Unexpected blues songs.\n")
    }
    tags.Del("baby please don't go", "blues")
    if len(tags.Tags) != 1 {
        t.Errorf("Wrong amount of tags.\n")
    }
}

func TestTagSaveToLoadFrom(t *testing.T) {
    tags := &Tags{}
    tags.Add("i'm in the band", "rOCk")
    tags.Add("red house", "RocK", "Blues")
    tags.Add("soul serenade", "soul")
    tags.Add("baby please don't go", "blues")
    tags.Add("feeling good", "jaZz")
    songs := tags.Get("rock")
    if len(songs) != 2 {
        t.Errorf("Wrong amount of rock songs.")
    } else if songs[0] != "i'm in the band" ||
              songs[1] != "red house" {
        t.Errorf("Unexpected rock songs.\n")
    }
    songs = tags.Get("blues")
    if len(songs) != 2 {
        t.Errorf("Wrong amount of blues songs.\n")
    } else if songs[0] != "red house" ||
              songs[1] != "baby please don't go" {
        t.Errorf("Unexpected blues songs.\n")
    }
    songs = tags.Get("soul")
    if len(songs) != 1 {
        t.Errorf("Wrong amount of soul songs.\n")
    } else if songs[0] != "soul serenade" {
        t.Errorf("Unexpected soul song.\n")
    }
    songs = tags.Get("jazz")
    if len(songs) != 1 {
        t.Errorf("Wrong amount of jazz songs.\n")
    } else if songs[0] != "feeling good" {
        t.Errorf("Unexpected jazz song.\n")
    }
    songs = tags.Get("BregaSofrenciaExtremo")
    if len(songs) != 0 {
        t.Errorf("Unexistent tag has returned songs!!!\n")
    }
    err := tags.SaveTo("tags.json")
    if err == nil {
        defer os.Remove("tags.json")
        newTags := &Tags{}
        err = newTags.LoadFrom("tags.json")
        songs := newTags.Get("rock")
        if len(songs) != 2 {
            t.Errorf("Wrong amount of rock songs.")
        } else if songs[0] != "i'm in the band" ||
                  songs[1] != "red house" {
            t.Errorf("Unexpected rock songs.\n")
        }
        songs = newTags.Get("blues")
        if len(songs) != 2 {
            t.Errorf("Wrong amount of blues songs.\n")
        } else if songs[0] != "red house" ||
                  songs[1] != "baby please don't go" {
            t.Errorf("Unexpected blues songs.\n")
        }
        songs = newTags.Get("soul")
        if len(songs) != 1 {
            t.Errorf("Wrong amount of soul songs.\n")
        } else if songs[0] != "soul serenade" {
            t.Errorf("Unexpected soul song.\n")
        }
        songs = newTags.Get("jazz")
        if len(songs) != 1 {
            t.Errorf("Wrong amount of jazz songs.\n")
        } else if songs[0] != "feeling good" {
            t.Errorf("Unexpected jazz song.\n")
        }
        songs = newTags.Get("BregaSofrenciaExtremo")
        if len(songs) != 0 {
            t.Errorf("Unexistent tag has returned songs!!!\n")
        }
    } else {
        t.Errorf("SaveTo() has failed : '%s'.\n", err.Error())
    }
}