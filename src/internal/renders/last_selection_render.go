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

func LastSelectionRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    templatedOutput := strings.Replace(templatedInput, vars.EutherpeTemplateNeedleLastSelection, eutherpeVars.LastSelection, -1)
    eutherpeVars.LastSelection = ""
    return templatedOutput
}
