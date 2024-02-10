package renders

import (
    "internal/vars"
    "fmt"
    "testing"
)

func TestShuffleModeRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleShuffleMode)
    output := ShuffleModeRender(templatedInput, eutherpeVars)
    if output != "Shuffle" {
        t.Errorf("ShuffleModeRender() seems not to be rendering accordingly.\n")
    } else {
        eutherpeVars.Player.Shuffle = true
        output = ShuffleModeRender(templatedInput, eutherpeVars)
        if output != "Original" {
            t.Errorf("ShuffleModeRender() seems not to be rendering accordingly.\n")
        }
    }
}
