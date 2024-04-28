package renders

import (
    "internal/vars"
    "strings"
)

func HostNameRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleHostName, eutherpeVars.HostName, -1)
}
