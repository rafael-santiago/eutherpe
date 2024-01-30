package renders

import (
    "internal/vars"
    "fmt"
    "testing"
)

func TestEutherpeAddrRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.HTTPd.Addr = "127.0.0.1:8080"
    templatedInput := fmt.Sprintf("http://%s/index.html", vars.EutherpeTemplateNeedleEutherpeAddr)
    output := EutherpeAddrRender(templatedInput, eutherpeVars)
    if output != "http://127.0.0.1:8080/index.html" {
        t.Errorf("EutherpeAddrRender() seems not to be working accordingly.\n")
    }
}
