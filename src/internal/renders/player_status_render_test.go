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
    if output != "{\"now-playing\":\"Queens Of The Stone Age - The Bronze\",\"album-cover-src\" : \"data:image/;base64,QmxhdQ==\"}" {
        t.Errorf("PlayerStatusRender() is not rendering accordingly : '%s'\n", output)
    }
}
