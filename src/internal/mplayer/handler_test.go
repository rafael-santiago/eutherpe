//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package mplayer

import (
    "testing"
)

func TestPlayStopDynamics(t *testing.T) {
    handle, err := Play("song.mp3", "../mplayer")
    if err != nil {
        t.Errorf("Play() has returned an error when it must not : '%s'.\n", err.Error())
    }
    err = Stop(handle)
    if err != nil {
        t.Errorf("Stop() has returned an error when it must not : '%s'.\n", err.Error())
    }
}

func TestGetVolumeLevel(t *testing.T) {
    volLevel := GetVolumeLevel("../mplayer")
    if  volLevel != 93 {
        t.Errorf("GetVolumeLevel() is returning wrong value : %d\n", volLevel)
    }
}

func TestSetVolume(t *testing.T) {
    SetVolume(10, "../mplayer")
}