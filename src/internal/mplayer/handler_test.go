package mplayer

import (
    "testing"
)

func TestPlayStopDynamics(t *testing.T) {
    handle, err := Play("song.mp3", "./")
    if err != nil {
        t.Errorf("Play() has returned an error when it must not : '%s'.\n", err.Error())
    }
    err = Stop(handle)
    if err != nil {
        t.Errorf("Stop() has returned an error when it must not : '%s'.\n", err.Error())
    }
}
