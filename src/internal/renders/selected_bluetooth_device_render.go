package renders

import (
    "internal/vars"
    "strings"
)

func SelectedBluetoothDeviceRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var selectedBluetoothDevice string
    if len(eutherpeVars.CachedDevices.BlueDevId) > 0 {
        selectedBluetoothDevice = eutherpeVars.CachedDevices.BlueDevId
    } else {
        selectedBluetoothDevice = "(null)"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleSelectedBluetoothDevice, selectedBluetoothDevice, -1)
}