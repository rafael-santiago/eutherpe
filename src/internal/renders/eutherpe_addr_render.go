package renders

import (
    "internal/vars"
    "strings"
)

func EutherpeAddrRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleEutherpeAddr, eutherpeVars.HTTPd.Addr, -1)
}
