//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package system

import (
    "os/exec"
    "path"
    "flag"
)

func Shutdown() error {
    return doShutdown("-h", "now")
}

func Reboot() error {
    return doShutdown("-r", "now")
}

func doShutdown(params ...string) error {
    var customPath string
    var app string
    if flag.Lookup("test.v") != nil {
        customPath = "../system"
        app = path.Join(customPath, "shutdown")
    } else {
        app = "sudo"
    }
    finalParams := make([]string, 0)
    if app == "sudo" {
        finalParams = append(finalParams, "/usr/sbin/shutdown")
    }
    for _, param := range params {
        finalParams = append(finalParams, param)
    }
    return exec.Command(app, finalParams...).Run()
}
