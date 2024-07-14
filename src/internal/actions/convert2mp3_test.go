//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package actions

import (
    "internal/vars"
    "testing"
    "os"
)

func TestConvert2MP3(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    err := Convert2MP3(eutherpeVars, nil)
    if err == nil {
        t.Errorf("Convert2MP3() has not failed when it should.\n")
    } else if err.Error() != "You need to set a storage device first." {
        t.Errorf("Convert2MP3() has failed with unexpected error.\n")
    } else {
        _ = os.Mkdir("/tmp/test", 0777)
        _ = os.WriteFile("/tmp/abc.m4a", []byte("abc"), 0777)
        _ = os.WriteFile("/tmp/test/yyz.mp4", []byte("yyz"), 0777)
        os.Remove("/tmp/abc.mp3")
        os.Remove("/tmp/test/yyz.mp3")
        defer os.Remove("/tmp/abc.m4a")
        defer os.Remove("/tmp/abc.mp3")
        defer os.RemoveAll("/tmp/test/")
        eutherpeVars.CachedDevices.MusicDevId = "/tmp"
        err = Convert2MP3(eutherpeVars, nil)
        if err != nil {
            t.Errorf("Convert2MP3() has failed when it should not.\n")
        } else {
            _, err = os.Stat("/tmp/abc.mp3")
            if err != nil {
                t.Errorf("/tmp/abc.mp3 not found!\n")
            }
            _, err = os.Stat("/tmp/test/yyz.mp3")
            if err != nil {
                t.Errorf("/tmp/test/yyz.mp3 not found!\n")
            }
        }
    }
}
