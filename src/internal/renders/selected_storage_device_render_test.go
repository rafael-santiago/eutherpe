package renders

import (
    "internal/vars"
    "testing"
    "fmt"
)

func TestSelectedStorageDeviceRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleSelectedStorageDevice)
    output := SelectedStorageDeviceRender(templatedInput, eutherpeVars)
    if output != "(null)" {
        t.Errorf("SelectedStorageDevice() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.CachedDevices.MusicDevId = "/dev/stordev42"
    output = SelectedStorageDeviceRender(templatedInput, eutherpeVars)
    if output != "/dev/stordev42" {
        t.Errorf("SelectedStorageDevice() seems not to be rendering accordingly.\n")
    }
}
