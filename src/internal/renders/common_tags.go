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
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleCommonTags, commonTagsHTML, -1)
}

