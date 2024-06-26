//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
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
