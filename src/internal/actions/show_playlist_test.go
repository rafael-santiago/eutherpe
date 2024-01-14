package actions

import (
    "internal/vars"
    "internal/dj"
    "net/url"
    "testing"
)

func TestShowPlaylist(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.RenderedPlaylist = "Top-42"
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist{ Name: "Ziriguidum!" })
    userData := &url.Values{}
    err := ShowPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ShowPlaylist() did not return error when it should.\n")
    } else if err.Error() != "Malformed playlist-show request." {
        t.Errorf("ShowPlaylist() has returned an unexpected error: '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "PediuPraTocarToco!")
    err = ShowPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ShowPlaylist() did not return error when it should.\n")
    } else if err.Error() != "Playlist 'PediuPraTocarToco!' not exists." {
        t.Errorf("ShowPlaylist() has returned an unexpected error: '%v'\n", err.Error())
    }
    if eutherpeVars.RenderedPlaylist != "Top-42" {
        t.Errorf("ShowPlaylist() has changed RenderedPlaylist register when it should not.\n")
    }
    userData.Del(vars.EutherpePostFieldPlaylist)
    userData.Add(vars.EutherpePostFieldPlaylist, "Ziriguidum!")
    err = ShowPlaylist(eutherpeVars, userData)
    if eutherpeVars.RenderedPlaylist != "Ziriguidum!" {
        t.Errorf("ShowPlaylist() has not changed RenderedPlaylist register as it should.\n")
    }
}
