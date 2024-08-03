//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package auth

import (
    "testing"
)

func TestHashKey(t *testing.T) {
    hashData := HashKey("123mudar*")
    if len(hashData) == 0 {
        t.Errorf("HashKey() seems not to be accordingly work.\n")
    }
}

func TestValidate(t *testing.T) {
    hashData := HashKey("123mudar*")
    if Validate("123mudar*", hashData[:4]) {
        t.Errorf("Validate() has succeeded when it should not.\n")
    }
    if !Validate("123mudar*", HashKey("123mudar*")) {
        t.Errorf("Validate() has not succeeded when it should.\n")
    }
    if Validate("123mudou", HashKey("123mudar*")) {
        t.Errorf("Validate() has succeeded when it should not.\n")
    }
}
