package renders

import (
    "internal/vars"
    "fmt"
    "testing"
)

func TestFoundStorageDevicesRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.StorageDevices = append(eutherpeVars.StorageDevices, "/dev/stordev42", "/media/rs/musicas", "/media/usbdrive101")
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleFoundStorageDevices)
    output := FoundStorageDevicesRender(templatedInput, eutherpeVars)
    if output != "<ul class=\"nested\"><input type=\"checkbox\" id=\"/dev/stordev42\" class=\"StorageDevice\" onclick=\"selectSingleElement(this);\">/dev/stordev42<br><input type=\"checkbox\" id=\"/media/rs/musicas\" class=\"StorageDevice\" onclick=\"selectSingleElement(this);\">/media/rs/musicas<br><input type=\"checkbox\" id=\"/media/usbdrive101\" class=\"StorageDevice\" onclick=\"selectSingleElement(this);\">/media/usbdrive101<br></ul>" {
        t.Errorf("FoundStorageDevicesRender() seems not to be rendering accordingly.\n")
    }
}