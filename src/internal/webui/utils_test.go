//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package webui

import (
    "testing"
)

func TestGetMIMEType(t *testing.T) {
    type TestCtx struct {
        FilePath string
        Expected string
    }
    testVector := []TestCtx {
        TestCtx { "eutherpe.html", "text/html" },
        TestCtx { "eutherpe.js", "text/javascript" },
        TestCtx { "eutherpe.css", "text/css" },
        TestCtx { "eutherpe.txt", "text/plain" },
    }
    for _, test := range testVector {
        if GetMIMEType(test.FilePath) != test.Expected {
            t.Errorf("GetMIMEType() has returned wrong MIME type.\n")
        }
    }
}
