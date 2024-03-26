package actions

import (
    "internal/vars"
    "internal/auth"
    "net/url"
    "fmt"
)

func ChangePassphrase(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    password, has := (*userData)[vars.EutherpePostFieldPassword]
    if !has {
        return fmt.Errorf("Malformed settings-changepassphrase request.")
    }
    newPassword, has := (*userData)[vars.EutherpePostFieldNewPassword]
    if !has {
        return fmt.Errorf("Malformed settings-changepassphrase request.")
    }
    if !auth.Validate(password[0], eutherpeVars.HTTPd.HashKey) {
        return fmt.Errorf("Wrong passphrase!")
    }
    eutherpeVars.HTTPd.HashKey = auth.HashKey(newPassword[0])
    return nil
}
