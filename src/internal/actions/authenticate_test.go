package actions

import (
    "internal/vars"
    "internal/auth"
    "net/url"
    "time"
    "testing"
)

func TestAuthenticate(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.HTTPd.AuthWatchdog = auth.NewAuthWatchdog(time.Duration(1 * time.Minute))
    eutherpeVars.HTTPd.HashKey = auth.HashKey("123mudar*")
    userData := &url.Values{}
    err := Authenticate(eutherpeVars, userData)
    if err == nil {
        t.Errorf("Authenticate() did not return an error when it should.")
    } else if err.Error() != "Malformed authenticate request." {
        t.Errorf("Authenticate() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpeActionId, "blauInvalido")
    err = Authenticate(eutherpeVars, userData)
    if err == nil {
        t.Errorf("Authenticate() did not return an error when it should.")
    } else if err.Error() != "Malformed authenticate request." {
        t.Errorf("Authenticate() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpeActionId)
    userData.Add(vars.EutherpeActionId, vars.EutherpeAuthenticateId)
    err = Authenticate(eutherpeVars, userData)
    if err == nil {
        t.Errorf("Authenticate() did not return an error when it should.")
    } else if err.Error() != "Malformed authenticate request." {
        t.Errorf("Authenticate() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldRemoteAddr, "127.0.0.1:1024")
    err = Authenticate(eutherpeVars, userData)
    if err == nil {
        t.Errorf("Authenticate() did not return an error when it should.")
    } else if err.Error() != "Malformed authenticate request." {
        t.Errorf("Authenticate() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPassword, "42")
    err = Authenticate(eutherpeVars, userData)
    if err == nil {
        t.Errorf("Authenticate() did not return an error when it should.")
    } else if err.Error() != "Wrong passphrase!" {
        t.Errorf("Authenticate did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldPassword)
    userData.Add(vars.EutherpePostFieldPassword, "123mudar*")
    err = Authenticate(eutherpeVars, userData)
    if err != nil {
        t.Errorf("Authenticate did return an error when it should not.\n")
    }
}
