package actions

import (
    "internal/vars"
    "net/url"
    "testing"
)

func TestSetCurrentConfig(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    userData.Set(vars.EutherpePostFieldConfig, "blau")
    err := SetCurrentConfig(eutherpeVars, userData)
    if err == nil {
        t.Errorf("SetCurrentConfig() is not returning error when it should.")
    } else if err.Error() != "Unknown config value." {
        t.Errorf("SetCurrentConfig() is returning unexpected error.")
    }
    configs := []string { vars.EutherpeConfigMusic,
                          vars.EutherpeConfigCollection,
                          vars.EutherpeConfigPlaylists,
                          vars.EutherpeConfigStorage,
                          vars.EutherpeConfigBluetooth,
                          vars.EutherpeConfigSettings,
    }
    for _, config := range configs {
        userData.Set(vars.EutherpePostFieldConfig, config)
        err = SetCurrentConfig(eutherpeVars, userData)
        if err != nil {
            t.Errorf("SetCurrentConfig() is returning error when it should not.")
        } else if eutherpeVars.CurrentConfig != config {
            t.Errorf("SetCurrentConfig() is not setting CurrentConfig accordingly.")
        }
    }
}
