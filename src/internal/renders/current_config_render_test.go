package renders

import (
    "internal/vars"
    "testing"
)

func TestCurrentConfigRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.CurrentConfig = "UnitTest"
    output := CurrentConfigRender(vars.EutherpeTemplateNeedleCurrentConfig, eutherpeVars)
    if output != "UnitTest" {
        t.Errorf("CurrentConfigRender() is not rendering accordingly.\n")
    }
}
