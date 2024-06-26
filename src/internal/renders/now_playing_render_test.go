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
    "fmt"
    "testing"
)

func TestNowPlayingRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleNowPlaying)
    output := NowPlayingRender(templatedInput, eutherpeVars)
    if len(output) > 0 {
        t.Errorf("NowPlayingRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.NowPlaying = mplayer.SongInfo { Title: "Lawman", Artist: "Motorhead" }
    output = NowPlayingRender(templatedInput, eutherpeVars)
    if output != "Motorhead - Lawman" {
        t.Errorf("NowPlayingRender() seems not to be rendering accordingly.\n")
    }
}
