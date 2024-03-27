package renders

import (
    "internal/vars"
    "strings"
    "fmt"
)

func UpNextCountRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    return metaCountRender(templatedInput,
                           vars.EutherpeTemplateNeedleUpNextCount,
                           func() string {
                                return fmt.Sprintf("%d", len(eutherpeVars.Player.UpNext))
                           })
}

func FoundStorageDevicesCountRender(templatedInput string,
                                    eutherpeVars *vars.EutherpeVars) string {
    return metaCountRender(templatedInput,
                           vars.EutherpeTemplateNeedleFoundStorageDevicesCount,
                           func() string {
                                return fmt.Sprintf("%d", len(eutherpeVars.StorageDevices))
                           })
}

func FoundBluetoothDevicesCountRender(templatedInput string,
                                    eutherpeVars *vars.EutherpeVars) string {
    return metaCountRender(templatedInput,
                           vars.EutherpeTemplateNeedleFoundBluetoothDevicesCount,
                           func() string {
                                return fmt.Sprintf("%d", len(eutherpeVars.BluetoothDevices))
                           })
}

func metaCountRender(templatedInput,
                     templateNeedle string,
                     countClosure func () string) string {
    return strings.Replace(templatedInput,
                           templateNeedle,
                           countClosure(), -1)
}
