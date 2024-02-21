package renders

import (
    "internal/vars"
    "strings"
    "fmt"
)

func VolumeLevelRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    volumeLevel := fmt.Sprintf("%d", eutherpeVars.Player.VolumeLevel)
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleVolumeLevel, volumeLevel, -1)
}
