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

func EutherpeAddrRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var addr string
    if !eutherpeVars.HTTPd.RequestedByHostName ||
       len(eutherpeVars.HostName) == 0 {
        addr = eutherpeVars.HTTPd.Addr
    } else {
        addr = eutherpeVars.HostName
    }
    addr += ":" + eutherpeVars.HTTPd.Port
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleEutherpeAddr, addr, -1)
}
