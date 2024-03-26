package actions

import (
    "internal/vars"
    "testing"
)

func TestFlickAuthMode(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    err := FlickAuthMode(eutherpeVars, nil)
    if err != nil {
        t.Errorf("FlickAuthMode() is failing when it should not.\n")
    } else if !eutherpeVars.HTTPd.Authenticated {
        t.Errorf("FlickAuthMode() is not enabling authenticated mode.\n")
    }
    err = FlickAuthMode(eutherpeVars, nil)
    if err != nil {
        t.Errorf("FlickAuthMode() is failing when it should not.\n")
    } else if eutherpeVars.HTTPd.Authenticated {
        t.Errorf("FlickAuthMode() is not disabling authenticated mode.\n")
    }
}
