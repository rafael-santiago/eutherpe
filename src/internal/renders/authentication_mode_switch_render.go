package renders

import (
    "internal/vars"
    "strings"
    "fmt"
)

func AuthenticationModeSwitchRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var authenticationModeSwitchHTML = "<input type=\"checkbox\" onclick=\"flickAuthenticationModeSwitch();\"%s> Ask passphrase"
    if eutherpeVars.HTTPd.Authenticated {
        authenticationModeSwitchHTML = fmt.Sprintf(authenticationModeSwitchHTML, " checked")
    } else {
        authenticationModeSwitchHTML = fmt.Sprintf(authenticationModeSwitchHTML, "")
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleAuthenticationModeSwitch, authenticationModeSwitchHTML, -1)
}