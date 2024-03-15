package actions

import (
    "internal/vars"
    "internal/mplayer"
    "internal/dj"
    "net/url"
    "testing"
)

func TestReproducePlaylist(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Collection = make(mplayer.MusicCollection)
    eutherpeVars.Collection["Sonic Youth"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Sonic Youth"]["Washing Machine"] = []mplayer.SongInfo {
        mplayer.SongInfo { "diamond-sea.mp3", "Diamond Sea", "Sonic Youth", "Washing Machine", "11", "1994", "", "Indie", },
    }
    eutherpeVars.Collection["Nirvana"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Nirvana"]["Unplugged"] = []mplayer.SongInfo {
        mplayer.SongInfo { "about-a-girl.mp3", "About a Girl", "Nirvana", "Unplugged", "01", "1993", "", "Grunge", },
    }
    eutherpeVars.Tags.Add("Sonic Youth/Washing Machine/Diamond Sea:diamond-sea.mp3", "90s", "Indie")
    eutherpeVars.Tags.Add("Nirvana/Unplugged/About a Girl:about-a-girl.mp3", "90s", "Grunge")
    userData := &url.Values{}
    err := ReproducePlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ReproducePlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Malformed playlist-reproduce request." {
        t.Errorf("PlayByGivenTags() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "LousyHits")
    err = ReproducePlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ReproducePlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Playlist 'LousyHits' not found!" {
        t.Errorf("ReproducePlaylist() did return an unexpected error : '%s'\n", err.Error())
    }
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist { Name: "IndieTunes" })
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Sonic Youth"]["Washing Machine"][0])
    userData.Del(vars.EutherpePostFieldPlaylist)
    userData.Add(vars.EutherpePostFieldPlaylist, "IndieTunes")
    err = ReproducePlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ReproducePlaylist() did return an error when it should not.\n")
    } else if eutherpeVars.Player.NowPlaying.FilePath != "diamond-sea.mp3" {
        t.Errorf("ReproducePlaylist() did not reproduce the expected song.\n")
    }
    MusicStop(eutherpeVars, nil)
}
