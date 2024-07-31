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
    "runtime"
)

func EutherpeRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    if len(eutherpeVars.APPName) == 0 && !strings.HasPrefix(runtime.GOARCH, "arm") {
        eutherpeVars.APPName = "Eutherpe"
    } else if len(eutherpeVars.APPName) == 0 {
        eutherpeVars.APPName = "Euther-&Pi;"
    }
    return strings.Replace(templatedInput,
                           vars.EutherpeTemplateNeedleEutherpe,
                           eutherpeVars.APPName, 1)
}
