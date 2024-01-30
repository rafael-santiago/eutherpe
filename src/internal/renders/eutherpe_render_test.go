package renders

import (
    "internal/vars"
    "testing"
    "fmt"
)

func TestEutherpeRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("data: %s", vars.EutherpeTemplateNeedleEutherpe)
    output := EutherpeRender(templatedInput, eutherpeVars)
    if eutherpeVars.APPName != "Eutherpe" {
        t.Errorf("APPName has unexpected value: '%s'\n", eutherpeVars.APPName)
    } else if output != "data: Eutherpe" {
        t.Errorf("EutherpeRender() seems not to be working accordingly.\n")
    }
}
