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
    "os"
)

func TestPlayStopDynamics(t *testing.T) {
    handle, err := Play("song.mp3", false, "../mplayer")
    if err != nil {
        t.Errorf("Play() has returned an error when it must not : '%s'.\n", err.Error())
    }
    err = Stop(handle)
    if err != nil {
        t.Errorf("Stop() has returned an error when it must not : '%s'.\n", err.Error())
    }
}

func TestGetVolumeLevel(t *testing.T) {
    volLevel := GetVolumeLevel(false, "../mplayer")
    if  volLevel != 93 {
        t.Errorf("GetVolumeLevel() is returning wrong value : %d\n", volLevel)
    }
}

func TestSetVolume(t *testing.T) {
    SetVolume(10, "", "../mplayer")
}

func TestConverToMP3MustPass(t *testing.T) {
    os.Remove("/tmp/foo.mp3")
    err := ConvertToMP3("/tmp/foo.m4a", "../mplayer")
    defer os.Remove("/tmp/foo.mp3")
    if err != nil {
        t.Errorf("ConvertToMP3() is failing when it should pass.\n")
    } else {
        _, err = os.Stat("/tmp/foo.mp3")
        if err != nil {
            t.Errorf("/tmp/foo.mp3 not found!\n")
        }
    }
}

func TestConvertToMP3MustFail(t *testing.T) {
    os.Setenv("FFMPEG_MUST_FAIL", "1")
    defer os.Unsetenv("FFMPEG_MUST_FAIL")
    os.Remove("/tmp/foo.mp3")
    err := ConvertToMP3("/tmp/foo.m4a", "../mplayer")
    defer os.Remove("/tmp/foo.mp3")
    if err == nil {
        t.Errorf("ConvertToMP3() is not failing when it should.\n")
    } else {
        _, err = os.Stat("/tmp/foo.mp3")
        if err == nil {
            t.Errorf("/tmp/foo.mp3 was found!\n")
        }
    }
}