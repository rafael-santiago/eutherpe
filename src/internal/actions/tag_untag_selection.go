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
    "strings"
    "fmt"
)

type TagEditionFunc func(filePath string, tags ...string)

func TagSelection(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    return metaTagSelection(eutherpeVars,
                            userData,
                            func(filePath string, tags ...string) {
                                eutherpeVars.Tags.Add(filePath, tags...)
                            })
}

func UntagSelection(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    return metaTagSelection(eutherpeVars,
                            userData,
                            func(filePath string, tags ...string) {
                                eutherpeVars.Tags.Del(filePath, tags...)
                            })
}

func metaTagSelection(eutherpeVars *vars.EutherpeVars, userData *url.Values, doTagEdition TagEditionFunc) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    selectionJSON, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed collection-tagselectionas/collection-untagselections request.")
    }
    rawTags, has := (*userData)[vars.EutherpePostFieldTags]
    if !has {
        return fmt.Errorf("Malformed collection-tagselectionas/collection-untagselections request.")
    }
    selection := ParseSelection(selectionJSON[0])
    tags := strings.Split(rawTags[0], ",")
    metaTagEdition(selection, tags, doTagEdition)
    eutherpeVars.SaveTags()
    return nil
}

func metaTagEdition(selection, tags []string, doTagEdition TagEditionFunc) {
    for _, selectionId := range selection {
        doTagEdition(selectionId, tags...)
    }
}
