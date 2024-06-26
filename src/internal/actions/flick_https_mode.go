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
    "syscall"
    "flag"
)

func FlickHTTPSMode(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    if flag.Lookup("test.v") == nil {
        defer syscall.Kill(syscall.Getpid(), syscall.SIGINT)
    }
    defer eutherpeVars.Unlock()
    eutherpeVars.HTTPd.TLS = !eutherpeVars.HTTPd.TLS
    return nil
}
