package renders

import (
    "internal/vars"
    "strings"
    "runtime"
)

func EutherpeRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    if len(eutherpeVars.APPName) == 0 && !strings.HasPrefix(runtime.GOARCH, "arm") {
        eutherpeVars.APPName = "Eutherpe"
    } else if len(eutherpeVars.APPName) == 0 {
        eutherpeVars.APPName = "Euther&Pi"
    }
    return strings.Replace(templatedInput,
                           vars.EutherpeTemplateNeedleEutherpe,
                           eutherpeVars.APPName, -1)
}
