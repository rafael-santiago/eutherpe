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
