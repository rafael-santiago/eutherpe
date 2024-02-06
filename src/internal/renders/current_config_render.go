package renders

import (
    "internal/vars"
    "strings"
)

func CurrentConfigRender(templatedString string, eutherpeVars *vars.EutherpeVars) string {
    return strings.Replace(templatedString, vars.EutherpeTemplateNeedleCurrentConfig, eutherpeVars.CurrentConfig, -1)
}
