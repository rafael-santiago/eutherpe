package renders

import (
    "internal/vars"
    "fmt"
    "testing"
)

func TestSelectedBluetoothDeviceRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleSelectedBluetoothDevice)
    output := SelectedBluetoothDeviceRender(templatedInput, eutherpeVars)
    if output != "(null)" {
        t.Errorf("SelectedBluetoothDeviceRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.CachedDevices.BlueDevId = "dirty-dog-blue"
    output = SelectedBluetoothDeviceRender(templatedInput, eutherpeVars)
    if output != "dirty-dog-blue" {
        t.Errorf("SelectedBluetoothDeviceRender() seems not to be rendering accordingly.\n")
    }
}
