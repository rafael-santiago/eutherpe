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
)

func LastErrorRender(templatedString string, eutherpeVars *vars.EutherpeVars) string {
    var errStr string
    if eutherpeVars.LastError != nil {
        errStr = eutherpeVars.LastError.Error()
    }
    return strings.Replace(templatedString, vars.EutherpeTemplateNeedleLastError, errStr, -1)
}
