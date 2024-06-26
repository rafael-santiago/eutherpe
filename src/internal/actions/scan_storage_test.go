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
    "net/url"
    "testing"
    "strings"
    "path"
    "io/ioutil"
    "os"
)

func TestScanStorage(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := ScanStorage(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ScanStorage() did not return an error when it should.\n")
    } else if err.Error() != "Unset MusicDevId." {
        t.Errorf("ScanStorage() did return an unexpected error.\n")
    }
    eutherpeVars.CachedDevices.MusicDevId = "/dev/42"
    err = ScanStorage(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ScanStorage() did not return an error when it should.\n")
    } else if err.Error() != "open /dev/42: no such file or directory" {
        t.Errorf("ScanStorage() did return an unexpected error.\n")
    }
    entries, err := os.ReadDir("../mplayer/test-data/")
    if err != nil {
        t.Errorf(err.Error())
    }
    for _, f := range entries {
        if strings.HasSuffix(f.Name(), ".id3") {
            destFilePath := path.Join(".", strings.Replace(f.Name(), ".id3", ".mp3", -1))
            data, _ := ioutil.ReadFile(path.Join("../mplayer/test-data", f.Name()))
            ioutil.WriteFile(destFilePath, data, 0644)
            defer os.Remove(destFilePath)
        }
    }
    eutherpeVars.CachedDevices.MusicDevId = "."
    err = ScanStorage(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ScanStorage() did return an error when it should not.\n")
    }
    if len(eutherpeVars.Collection) == 0 {
        t.Errorf("ScanStorage() seems not to be scanning the device properly.\n")
    }
    eutherpeVars.CachedDevices.MusicDevId = "/dev/ziriguidum"
    err = ScanStorage(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ScanStorage() did not return an error when it should.\n")
    } else if err.Error() != "open /dev/ziriguidum: no such file or directory" {
        t.Errorf("ScanStorage() did return an unexpected error.\n")
    } else if len(eutherpeVars.Collection) == 0 {
        t.Errorf("ScanStorage() seems to be clearing previous collection on errors.\n")
    }
}
