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
                               ".js": "text/javascript" }
}
