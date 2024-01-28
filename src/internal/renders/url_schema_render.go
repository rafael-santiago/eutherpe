package renders

import (
    "internal/vars"
    "strings"
)

func URLSchemaRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleURLSchema, eutherpeVars.HTTPd.URLSchema, -1)
}
