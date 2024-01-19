package actions

import (
    "internal/vars"
    "internal/bluebraces"
    "net/url"
    "flag"
    "fmt"
    "reflect"
    "runtime"
)

type unTrustFunc func(id string, customPath ...string) error

func TrustBluetoothDevice(_ *vars.EutherpeVars,
                          userData *url.Values) error {
    return unTrustMetaAction(userData, bluebraces.TrustDevice)
}

func UntrustBluetoothDevice(_ *vars.EutherpeVars,
                            userData *url.Values) error {
    return unTrustMetaAction(userData, bluebraces.UntrustDevice)
}

func unTrustMetaAction(userData *url.Values,
                       doUnTrust unTrustFunc) error {
    bluetoothDevice, has := (*userData)[vars.EutherpePostFieldBluetoothDevice]
    if !has {
        goLangHasNoTernaryAndItSucksABunch := []string {
            "bluetooth-untrust",
            "bluetooth-trust",
        }
        isTrustDevice := (runtime.FuncForPC(reflect.ValueOf(doUnTrust).Pointer()).Name() == "internal/bluebraces.TrustDevice")
        return fmt.Errorf("Malformed %s request.", goLangHasNoTernaryAndItSucksABunch[
                                                            // ZZZ(Rafael): No conversion between bool and int, too Hauhauahauah!
                                                            //              Back to kindergarten... gosh!
                                                            func () uint8 {
                                                                if isTrustDevice {
                                                                    return 1
                                                                }
                                                                return 0
                                                            }()])
    }
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    return doUnTrust(bluetoothDevice[0], customPath)
}
