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

func CopyrightRender(templatedInput string, _ *vars.EutherpeVars) string {
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleCopyrightDisclaimer,
                           vars.EutherpeCopyrightDisclaimer, -1)
}
