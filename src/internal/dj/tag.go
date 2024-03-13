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
    Tags []TagEntry
}

func (t *Tags) Get(tag string) []string {
    normalizedTag := strings.ToLower(tag)
    for _, currTag := range t.Tags {
        if normalizedTag == currTag.Name {
            return currTag.FilePaths
        }
    }
    return make([]string, 0)
}

func (t *Tags) Add(filePath string, tags ...string) {
    for _, tag := range tags {
        tagIndex := t.getTagIndex(tag)
        if tagIndex == -1 {
            t.Tags = append(t.Tags, TagEntry { strings.Trim(strings.ToLower(tag), " "), []string { filePath } })
        } else if !t.Tags[tagIndex].IsTaggedAlready(filePath) {
            t.Tags[tagIndex].FilePaths = append(t.Tags[tagIndex].FilePaths, filePath)
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
        t.Tags[tagIndex].Del(filePath)
        if len(t.Tags[tagIndex].FilePaths) == 0 {
            emptiedTags = append(emptiedTags, tag)
        }
    }
    for _, tag := range emptiedTags {
        for tagIndex, _ := range t.Tags {
            if t.Tags[tagIndex].Name == strings.Trim(tag, " ") {
                t.Tags = append(t.Tags[:tagIndex], t.Tags[(tagIndex + 1):]...)
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
    for _, currTag := range t.Tags {
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
    if len(t.Tags) == 0 {
        return -1
    }
    for p, currTag := range t.Tags {
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
