package actions

import (
    "internal/vars"
    "internal/auth"
    "net/url"
    "fmt"
)

func Authenticate(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    action, has := (*userData)[vars.EutherpeActionId]
    if !has || action[0] != vars.EutherpeAuthenticateId {
        return fmt.Errorf("Malformed authenticate request.")
    }
    remoteAddr, has := (*userData)[vars.EutherpePostFieldRemoteAddr]
    if !has {
        return fmt.Errorf("Malformed authenticate request.")
    }
    password, has := (*userData)[vars.EutherpePostFieldPassword]
    if !has {
        return fmt.Errorf("Malformed authenticate request.")
    }
    if !auth.Validate(password[0], eutherpeVars.HTTPd.HashKey) {
        return fmt.Errorf("Wrong passphrase!")
    }
    eutherpeVars.HTTPd.AuthWatchdog.RefreshAuthWindow(remoteAddr[0])
    return nil
}