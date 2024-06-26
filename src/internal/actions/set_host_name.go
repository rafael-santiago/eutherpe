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
    "internal/mdns"
    "net/url"
    "strings"
    "fmt"
)

func SetHostName(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    hostname, has := (*userData)[vars.EutherpePostFieldHostName]
    if !has {
        return fmt.Errorf("Malformed settings-sethostname request.")
    }
    if !strings.HasSuffix(hostname[0], ".local") {
        hostname[0] += ".local"
    }
    eutherpeVars.HostName = hostname[0]
    eutherpeVars.MDNS.GoinHome <- true
    eutherpeVars.MDNS.Hosts[0].Name = hostname[0]
    go mdns.MDNSServerStart(eutherpeVars.MDNS.Hosts, eutherpeVars.MDNS.GoinHome)
    return eutherpeVars.SaveSession()
}