//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "internal/mplayer"
    "internal/dj"
    "internal/bluebraces"
    _ "internal/storage"
    "os"
    "testing"
)

func TestRenderData(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Collection = make(mplayer.MusicCollection)
    eutherpeVars.Collection["Queens Of The Stone Age"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Motorhead"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["The Cramps"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"] = []mplayer.SongInfo {
        mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", },
    }
    eutherpeVars.Collection["Motorhead"]["Bomber"] = []mplayer.SongInfo {
        mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["Motorhead"]["Overkill"] = []mplayer.SongInfo {
        mplayer.SongInfo { "overkill.mp3", "Overkill", "Motorhead", "Overkill", "1", "1979", "", "Speed Metal", },
        mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
        mplayer.SongInfo { "pay-your-price.mp3", "(I Won't) Pay Your Price", "Motorhead", "Overkill", "3", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"] = []mplayer.SongInfo {
        mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    stoner := dj.Playlist { Name: "Stoner Rock" }
    stoner.Add(eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0])
    speedMetalAndPsychobilly := dj.Playlist { Name: "Speed Metal & Psychobilly" }
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["Motorhead"]["Overkill"][2])
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["Motorhead"]["Overkill"][0])
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["Motorhead"]["Overkill"][1])
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, stoner, speedMetalAndPsychobilly)
    eutherpeVars.StorageDevices = append(eutherpeVars.StorageDevices, "/dev/stordev42", "/media/rs/musicas", "/media/usbdrive101")
    eutherpeVars.CachedDevices.MusicDevId = "/media/rs/musicas"
    eutherpeVars.BluetoothDevices = append(eutherpeVars.BluetoothDevices, bluebraces.BluetoothDevice { Id: "/dev/blue42", Alias: "/dev/blue42" },
                                           bluebraces.BluetoothDevice { Id: "/dev/deepblue42", Alias: "/dev/deepblue42" },
                                           bluebraces.BluetoothDevice { Id: "/dev/dentucaumazulnofundodoseu___bolsuhhhh", Alias: "/dev/dentucaumazulnofundodoseu___bolsuhhhh"  })
    eutherpeVars.CachedDevices.BlueDevId = "/dev/deepblue42"
    eutherpeVars.Player.NowPlaying = eutherpeVars.Collection["Motorhead"]["Overkill"][1]
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["Motorhead"]["Overkill"][1])
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0])
    eutherpeVars.HTTPd.URLSchema = "http"
    eutherpeVars.HTTPd.Addr = "127.0.0.1:80"

    eutherpeHTML, err:= os.ReadFile("../../web/html/eutherpe.html")

    if err != nil {
        t.Errorf("Error when trying to read 'eutherpe.html'.\n")
    }

    output := RenderData(string(eutherpeHTML), eutherpeVars)
    expectedEutherpeHTML, err := os.ReadFile("test-data/expected-eutherpe.html")
    if err != nil {
        t.Errorf("Error when trying to read 'expected-eutherpe.html'.\n")
    }
    if output != string(expectedEutherpeHTML) {
        t.Errorf("RenderData() seems not to be rendering accordingly.\n")
    }
}
