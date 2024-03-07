package actions

import (
    "internal/vars"
    "internal/dj"
    "net/url"
    "fmt"
)

func GetCommonTags(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    selectionJSON, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed get-commontags request.")
    }
    selection := ParseSelection(selectionJSON[0])
    eutherpeVars.LastCommonTags = getCommonTagsFromSelection(selection, eutherpeVars.Tags)
    eutherpeVars.LastSelection = getLastSelection(selection)
    if len(eutherpeVars.LastCommonTags) > 0 {
        eutherpeVars.CurrentConfig = "RemoveTags"
    }
    return nil
}

func getCommonTagsFromSelection(selection []string, tagsRepo dj.Tags) []string {
    tagsRefCount := make(map[string]int, 0)
    for _, currSelection := range selection {
        filePath := GetSongFilePathFromSelectionId(currSelection)
        tagsFromFile := tagsRepo.GetTagsFromFile(filePath)
        for _, currTag := range tagsFromFile {
            tagsRefCount[currTag] += 1
        }
    }
    tags := make([]string, 0)
    for tag, count := range tagsRefCount {
        if len(selection) == 1 || count > 1 {
            tags = append(tags, tag)
        }
    }
    return tags
}

func getLastSelection(selection []string) string {
    var lastSelection string
    for c, currSelection := range selection {
        lastSelection += GetSongFilePathFromSelectionId(currSelection)
        if (c + 1) != len(selection) {
            lastSelection += ","
        }
    }
    return lastSelection
}
