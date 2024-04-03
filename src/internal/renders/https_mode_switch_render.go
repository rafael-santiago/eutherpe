package renders

import (
    "internal/vars"
    "strings"
    "fmt"
)

func HTTPSModeSwitchRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var HTTPSModeSwitchHTML = "<input type=\"checkbox\" onclick=\"flickHTTPSModeSwitch();\"%s> HTTPS"
    if eutherpeVars.HTTPd.TLS {
        HTTPSModeSwitchHTML = fmt.Sprintf(HTTPSModeSwitchHTML, " checked")
    } else {
        HTTPSModeSwitchHTML = fmt.Sprintf(HTTPSModeSwitchHTML, "")
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleHTTPSModeSwitch, HTTPSModeSwitchHTML, -1)
}