package renders

import (
    "internal/vars"
    "strings"
)

func EutherpeAddrRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var addr string
    if len(eutherpeVars.HostName) == 0 {
        addr = eutherpeVars.HTTPd.Addr
    } else {
        addr = eutherpeVars.HostName
    }
    addr += ":" + eutherpeVars.HTTPd.Port
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleEutherpeAddr, addr, -1)
}
