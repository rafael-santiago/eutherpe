package renders

import (
    "internal/vars"
    "testing"
)

func TestCommonTagsRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.LastCommonTags = make([]string, 0)
    eutherpeVars.LastCommonTags = append(eutherpeVars.LastCommonTags, "xablau", "blau-do-borogodoh", "eu-sei-sou-maluco")
    output := CommonTagsRender(vars.EutherpeTemplateNeedleCommonTags, eutherpeVars)
    if output != "<ul class=\"nested\"><input type=\"checkbox\" id=\"xablau\" class=\"Tag\" checked>xablau<br><input type=\"checkbox\" id=\"blau-do-borogodoh\" class=\"Tag\" checked>blau-do-borogodoh<br><input type=\"checkbox\" id=\"eu-sei-sou-maluco\" class=\"Tag\" checked>eu-sei-sou-maluco<br></ul>" {
        t.Errorf("CommonTagsRender() is not rendering accordingly : %s\n", output)
    }
}
