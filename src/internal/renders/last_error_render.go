package renders

import (
    "internal/vars"
    "strings"
)

func LastErrorRender(templatedString string, eutherpeVars *vars.EutherpeVars) string {
    var errStr string
    if eutherpeVars.LastError != nil {
        errStr = eutherpeVars.LastError.Error()
    }
    return strings.Replace(templatedString, vars.EutherpeTemplateNeedleLastError, errStr, -1)
}
