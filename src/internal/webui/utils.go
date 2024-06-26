//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package webui

import (
    "strings"
    "path/filepath"
)

func GetMIMEType(filePath string) string {
    mType := supportedMIMETypes()[strings.ToLower(filepath.Ext(filePath))]
    if len(mType) == 0 {
        mType = "text/plain"
    }
    return mType
}

func supportedMIMETypes() map[string]string {
    return map[string]string { ".html": "text/html",
                               ".css": "text/css",
                               ".js": "text/javascript",
                               ".cer": "application/pkix-cert", }
}
