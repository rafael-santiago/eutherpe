package renders

import (
    "internal/vars"
    "internal/bluebraces"
    "fmt"
    "testing"
)

func TestFoundBluetoothDevices(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.BluetoothDevices = append(eutherpeVars.BluetoothDevices, bluebraces.BluetoothDevice { Id: "/dev/blue42", Alias: "/dev/blue42" },
                                           bluebraces.BluetoothDevice { Id: "/dev/deepblue42", Alias: "/dev/deepblue42" },
                                           bluebraces.BluetoothDevice { Id: "/dev/dentucaumazulnofundodoseu___bolsuhhhh", Alias: "/dev/dentucaumazulnofundodoseu___bolsuhhhh"  })
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleFoundBluetoothDevices)
    output := FoundBluetoothDevicesRender(templatedInput, eutherpeVars)
    if output != "<ul class=\"nested\"><input type=\"checkbox\" id=\"/dev/blue42\" class=\"BluetoothDevice\" onclick=\"selectSingleElement(this);\">/dev/blue42<br><input type=\"checkbox\" id=\"/dev/deepblue42\" class=\"BluetoothDevice\" onclick=\"selectSingleElement(this);\">/dev/deepblue42<br><input type=\"checkbox\" id=\"/dev/dentucaumazulnofundodoseu___bolsuhhhh\" class=\"BluetoothDevice\" onclick=\"selectSingleElement(this);\">/dev/dentucaumazulnofundodoseu___bolsuhhhh<br></ul>" {
        t.Errorf("FoundBluetoothDevices() seems not to be rendering accordingly.\n")
    }
}
