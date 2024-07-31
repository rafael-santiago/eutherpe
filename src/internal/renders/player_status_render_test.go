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
    "testing"
)

func TestPlayerStatusRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Player.NowPlaying = mplayer.SongInfo { "the-bronze.mp3", "The Bronze", "Queens Of The Stone Age", "Queens Of The Stone Age", "6", "1998", "Blau", "Stoner Rock", }
    output := PlayerStatusRender(vars.EutherpeTemplateNeedlePlayerStatus, eutherpeVars)
    if output != "{\"now-playing\":\"Queens Of The Stone Age - The Bronze\",\"album-cover-src\" : \"data:image/gif;base64,R0lGODlhKEdJRiBHT0VTIEhFUkUuLi4p\"}" {
        t.Errorf("PlayerStatusRender() is not rendering accordingly : '%s'\n", output)
    }
}
