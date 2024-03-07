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
        return fmt.Errorf("Malformed collection-un/tagselectionas/collection-untagselections request.")
    }
    selection := ParseSelection(selectionJSON[0])
    tags := strings.Split(rawTags[0], ",")
    metaTagEdition(selection, tags, doTagEdition)
    return nil
}

func metaTagEdition(selection, tags []string, doTagEdition TagEditionFunc) {
    for _, selectionId := range selection {
        filePath := GetSongFilePathFromSelectionId(selectionId)
        if len(filePath) == 0 {
            filePath = selectionId
        }
        doTagEdition(filePath, tags...)
    }
}
