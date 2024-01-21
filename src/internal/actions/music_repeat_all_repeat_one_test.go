package actions

import (
    "internal/vars"
    "net/url"
    "testing"
)

func TestMusicRepeatAll(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    eutherpeVars.Player.RepeatOne = true
    err := MusicRepeatAll(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicRepeatAll() has returned an error when it should not.\n")
    } else if !eutherpeVars.Player.RepeatAll {
        t.Errorf("MusicRepeatAll() seems not to be setting the RepeatAll flag accordingly.\n")
    } else if eutherpeVars.Player.RepeatOne {
        t.Errorf("MusicRepeatAll() seems not to be unsetting the RepeatOne flag accordingly.\n")
    } else if MusicRepeatAll(eutherpeVars, userData) != nil {
        t.Errorf("MusicRepeatAll() has returned an error when it should not.\n")
    } else if eutherpeVars.Player.RepeatAll {
        t.Errorf("MusicRepeatAll() seems not to be unsetting the RepeatAll flag accordingly.\n")
    } else if eutherpeVars.Player.RepeatOne {
        t.Errorf("MusicRepeatAll() seems to be setting the RepeatOne flag.\n")
    }
}

func TestMusicRepeatOne(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    eutherpeVars.Player.RepeatAll = true
    err := MusicRepeatOne(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicRepeatOne() has returned an error when it should not.\n")
    } else if !eutherpeVars.Player.RepeatOne {
        t.Errorf("MusicRepeatOne() seems not to be setting the RepeatOne flag accordingly.\n")
    } else if eutherpeVars.Player.RepeatAll {
        t.Errorf("MusicRepeatOne() seems not to be unsetting the RepeatAll flag accordingly.\n")
    } else if MusicRepeatOne(eutherpeVars, userData) != nil {
        t.Errorf("MusicRepeatOne() has returned an error when it should not.\n")
    } else if eutherpeVars.Player.RepeatOne {
        t.Errorf("MusicRepeatOne() seems not to be unsetting the RepeatOne flag accordingly.\n")
    } else if eutherpeVars.Player.RepeatAll {
        t.Errorf("MusicRepeatAll() seems to be setting the RepeatAll flag.\n")
    }
}