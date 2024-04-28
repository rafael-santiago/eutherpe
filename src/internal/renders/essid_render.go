package renders

import (
    "internal/vars"
    "strings"
)

func ESSIDRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleESSID, eutherpeVars.WLAN.ESSID, -1)
}
