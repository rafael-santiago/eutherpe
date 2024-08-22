//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package options

import (
    "testing"
    "os"
)

func TestGet(t *testing.T) {
    os.Args = append(os.Args, "--testGet=42")
    value := Get("--testGet")
    if len(value) != 0 {
        t.Errorf("Get() is not returning empty string.\n")
    }
    value = Get("testGet")
    if value != "42" {
        t.Errorf("Get() is not returning the expected value.\n")
    }
    value = Get("not-found", "404")
    if value != "404" {
        t.Errorf("Get() is not returning the default value.\n")
    }
}

func TestHasFlag(t *testing.T) {
    os.Args = append(os.Args, "--testHasFlag")
    if HasFlag("--testHasFlag") {
        t.Errorf("HasFlag() is not returning false as expected.\n")
    }
    if !HasFlag("testHasFlag") {
        t.Errorf("HasFlag() is not returning true as expected.\n")
    }
}
