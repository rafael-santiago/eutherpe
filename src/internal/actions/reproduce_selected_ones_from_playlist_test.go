package actions

import (
    "internal/vars"
    "internal/mplayer"
    "internal/dj"
    "net/url"
    "testing"
)

func TestReproduceSelectedOnesFromPlaylist(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Player.UpNextCurrentOffset = 0
    eutherpeVars.Player.Stopped = true
    eutherpeVars.Collection = make(mplayer.MusicCollection)
    eutherpeVars.Collection["Sonic Youth"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Sonic Youth"]["Washing Machine"] = []mplayer.SongInfo {
        mplayer.SongInfo { "/diamond-sea.mp3", "Diamond Sea", "Sonic Youth", "Washing Machine", "11", "1994", "", "Indie", },
    }
    eutherpeVars.Collection["Nirvana"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Nirvana"]["Unplugged"] = []mplayer.SongInfo {
        mplayer.SongInfo { "/about-a-girl.mp3", "About a Girl", "Nirvana", "Unplugged", "01", "1993", "", "Grunge", },
    }
    eutherpeVars.Tags.Add("Sonic Youth/Washing Machine/Diamond Sea:/diamond-sea.mp3", "90s", "Indie")
    eutherpeVars.Tags.Add("Nirvana/Unplugged/About a Girl:/about-a-girl.mp3", "90s", "Grunge")
    userData := &url.Values{}
    err := ReproduceSelectedOnesFromPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Malformed playlist-reproduceselectedones request." {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "LousyHits")
    err = ReproduceSelectedOnesFromPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Malformed playlist-reproduceselectedones request." {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[\"90s:Nirvana/Unplugged/:/about-a-girl.mp3\"]")
    err = ReproduceSelectedOnesFromPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Playlist 'LousyHits' has not found." {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did return an unexpected error.\n")
    }
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist { Name: "90s" })
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Sonic Youth"]["Washing Machine"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Nirvana"]["Unplugged"][0])
    userData.Del(vars.EutherpePostFieldPlaylist)
    userData.Add(vars.EutherpePostFieldPlaylist, "90s")
    err = ReproduceSelectedOnesFromPlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did return an error when it should not : '%s'\n", err.Error())
    } else if eutherpeVars.Player.NowPlaying.FilePath != "/about-a-girl.mp3" {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did not reproduce the expected song.\n")
    }
    if len(eutherpeVars.Player.UpNext) != 1 {
        t.Errorf("ReproduceSelectedOnesFromPlaylist() did not selected the exact amount of songs.\n")
    }
    MusicStop(eutherpeVars, nil)
}
