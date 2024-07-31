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

func CommonTagsRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    commonTagsHTML := "<ul class=\"nested\">"
    for _, currTag := range eutherpeVars.LastCommonTags {
        commonTagsHTML += "<input type=\"checkbox\" id=\"" + currTag + "\" class=\"Tag\" checked>" + currTag + "<br>"
    }
    commonTagsHTML += "</ul>"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleCommonTags, commonTagsHTML, 1)
}

