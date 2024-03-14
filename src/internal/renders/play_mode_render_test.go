package renders

import (
    "internal/vars"
    "internal/mplayer"
    "fmt"
    "testing"
)

func TestPlayModeRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedlePlayMode)
    output := PlayModeRender(templatedInput, eutherpeVars)
    if output != "&#x25BA" {
        t.Errorf("PlayModeRender() seems not to be rendering accordingly : %s\n", output)
    } else {
        eutherpeVars.Player.NowPlaying = mplayer.SongInfo { FilePath: "/dev/42/carnavoyeur.mp3" }
        output = PlayModeRender(templatedInput, eutherpeVars)
        if output != "&#x25A0" {
            t.Errorf("PlayModeRender() seems not to be rendering accordingly : %s\n", output)
        }
    }
}
