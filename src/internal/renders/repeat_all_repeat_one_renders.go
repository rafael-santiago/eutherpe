//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "strings"
    "fmt"
)

func RepeatAllRender(templatedString string, eutherpeVars *vars.EutherpeVars) string {
    var repeatAllHTML = "<input type=\"checkbox\" onclick=\"musicRepeatAll();\"%s><small>Repeat All</small>"
    return metaRepeatModeRender(templatedString, repeatAllHTML,
                                vars.EutherpeTemplateNeedleRepeatAll, eutherpeVars.Player.RepeatAll)
}

func RepeatOneRender(templatedString string, eutherpeVars *vars.EutherpeVars) string {
    var repeatAllHTML = "<input type=\"checkbox\" onclick=\"musicRepeatOne();\"%s><small>Repeat One</small>"
    return metaRepeatModeRender(templatedString, repeatAllHTML,
                                vars.EutherpeTemplateNeedleRepeatOne, eutherpeVars.Player.RepeatOne)
}

func metaRepeatModeRender(templatedString, innerHTML, templateNeedle string, statusFlag bool) string {
    return strings.Replace(templatedString, templateNeedle, fmt.Sprintf(innerHTML, func() string {
                                                                                     if statusFlag {
                                                                                        return " checked"
                                                                                     }
                                                                                     return ""
                                                                        }()), -1)
}

