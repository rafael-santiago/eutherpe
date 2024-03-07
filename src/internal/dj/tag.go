package dj

import (
    "encoding/json"
    "os"
    "strings"
)

type TagEntry struct {
    Name string
    FilePaths []string
}

type Tags struct {
    tags []TagEntry
}

func (t *Tags) Add(filePath string, tags ...string) {
    for _, tag := range tags {
        tagIndex := t.getTagIndex(tag)
        if tagIndex == -1 {
            t.tags = append(t.tags, TagEntry { strings.Trim(tag, " "), []string { filePath } })
        } else if !t.tags[tagIndex].IsTaggedAlready(filePath) {
            t.tags[tagIndex].FilePaths = append(t.tags[tagIndex].FilePaths, filePath)
        }
    }
}

func (t *Tags) Del(filePath string, tags ...string) {
    emptiedTags := make([]string, 0)
    for _, tag := range tags {
        tagIndex := t.getTagIndex(tag)
        if tagIndex == -1 {
            continue
        }
        t.tags[tagIndex].Del(filePath)
        if len(t.tags[tagIndex].FilePaths) == 0 {
            emptiedTags = append(emptiedTags, tag)
        }
    }
    for _, tag := range emptiedTags {
        for tagIndex, _ := range t.tags {
            if t.tags[tagIndex].Name == strings.Trim(tag, " ") {
                t.tags = append(t.tags[:tagIndex], t.tags[(tagIndex + 1):]...)
                break
            }
        }
    }
}

func (t *Tags) SaveTo(filePath string) error {
    data, err := json.Marshal(*t)
    if err != nil {
        return err
    }
    return os.WriteFile(filePath, data, 0777)
}

func (t *Tags) LoadFrom(filePath string) error {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, t)
}

func (t *Tags) GetTagsFromFile(filePath string) []string {
    tags := make([]string, 0)
    for _, currTag := range t.tags {
        for _, currFilePath := range currTag.FilePaths {
            if currFilePath == filePath {
                tags = append(tags, currTag.Name)
                break
            }
        }
    }
    return tags
}

func (t *Tags) getTagIndex(tagName string) int {
    if len(t.tags) == 0 {
        return -1
    }
    for p, currTag := range t.tags {
        if currTag.Name == tagName {
            return p
        }
    }
    return -1
}

func (t *TagEntry) IsTaggedAlready(filePath string) bool {
    if len(t.FilePaths) == 0 {
        return false
    }
    for _, currFilePath := range t.FilePaths {
        if currFilePath == filePath {
            return true
        }
    }
    return false
}

func (t *TagEntry) Del(filePath string) {
    for f, currFilePath := range t.FilePaths {
        if currFilePath == filePath {
            t.FilePaths = append(t.FilePaths[:f], t.FilePaths[(f+1):]...)
            break
        }
    }
}
