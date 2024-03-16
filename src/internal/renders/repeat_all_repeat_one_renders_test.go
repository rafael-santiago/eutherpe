package renders

import (
    "internal/vars"
    "testing"
)

func TestRepeatAllRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    output := RepeatAllRender(vars.EutherpeTemplateNeedleRepeatAll, eutherpeVars)
    if output != "<input type=\"checkbox\" onclick=\"musicRepeatAll();\"><small>Repeat All</small>" {
        t.Errorf("RepeatAllRender() is not rendering accordingly : '%s'\n", output)
    }
    eutherpeVars.Player.RepeatAll = true
    output = RepeatAllRender(vars.EutherpeTemplateNeedleRepeatAll, eutherpeVars)
    if output != "<input type=\"checkbox\" onclick=\"musicRepeatAll();\" checked><small>Repeat All</small>" {
        t.Errorf("RepeatAllRender() is not rendering accordingly : '%s'\n", output)
    }
}

func TestRepeatOneRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    output := RepeatOneRender(vars.EutherpeTemplateNeedleRepeatOne, eutherpeVars)
    if output != "<input type=\"checkbox\" onclick=\"musicRepeatOne();\"><small>Repeat One</small>" {
        t.Errorf("RepeatOneRender() is not rendering accordingly : '%s'\n", output)
    }
    eutherpeVars.Player.RepeatOne = true
    output = RepeatOneRender(vars.EutherpeTemplateNeedleRepeatOne, eutherpeVars)
    if output != "<input type=\"checkbox\" onclick=\"musicRepeatOne();\" checked><small>Repeat One</small>" {
        t.Errorf("RepeatOneRender() is not rendering accordingly : '%s'\n", output)
    }
}
