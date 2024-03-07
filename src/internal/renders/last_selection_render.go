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
