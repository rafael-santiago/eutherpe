package renders

import (
    "internal/vars"
    "testing"
)

func TestVolumeLevelRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Player.VolumeLevel = 42
    output := VolumeLevelRender(vars.EutherpeTemplateNeedleVolumeLevel, eutherpeVars)
    if output != "42" {
        t.Errorf("VolumeLevelRender() is not rendering accordingly : '%s'\n", output)
    }
}
