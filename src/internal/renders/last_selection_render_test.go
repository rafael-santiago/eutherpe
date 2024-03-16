package renders

import (
    "internal/vars"
    "testing"
)

func TestLastSelectionRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.LastSelection = "/Bugzinho/quando/nasce/quase/mela/todo/chao/"
    output := LastSelectionRender(vars.EutherpeTemplateNeedleLastSelection, eutherpeVars)
    if output != "/Bugzinho/quando/nasce/quase/mela/todo/chao/" {
        t.Errorf("LastSelectionRender() is not rendering accordingly.\n")
    }
    if len(eutherpeVars.LastSelection) > 0 {
        t.Errorf("LastSelectionRender() is not clearing LastSelection register.\n")
    }
}
