//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package options

import (
    "strings"
    "os"
)

func Get(option string, defaultValue ...string) string {
    temp := "--" + option + "="
    var optionDefault string
    if len(defaultValue) > 0 {
        optionDefault = defaultValue[0]
    }
    for _, arg := range os.Args {
        if strings.HasPrefix(arg, temp) {
            return arg[len(temp):]
        }
    }
    return optionDefault
}

func HasFlag(option string) bool {
    temp := "--" + option
    for _, arg := range os.Args {
        if arg == temp {
            return true
        }
    }
    return false
}
