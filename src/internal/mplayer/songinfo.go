package mplayer

import (
    "os"
    "fmt"
    "strings"
    "regexp"
    "path"
    "unicode/utf8"
)

const (
    kSongInfoFieldsNr = 8
)

type SongInfo struct {
    FilePath string
    Title string
    Artist string
    Album string
    TrackNumber string
    Year string
    AlbumCover string
    Genre string
}

func ScanSongs(basePath string) ([]SongInfo, error) {
    files, err := os.ReadDir(basePath)
    if err != nil {
        return make([]SongInfo, 0), err
    }
    songs := make([]SongInfo, 0)
    for _, file := range files {
        fileName := file.Name()
        if path.Ext(fileName) == ".mp3" || path.Ext(fileName) == ".mp4" {
            songInfo, err := GetSongInfo(path.Join(basePath, fileName))
            if err != nil {
                continue
            }
            songs = append(songs, songInfo)
        } else {
            filePath := path.Join(basePath, fileName)
            stat, err := os.Stat(filePath)
            if err == nil && stat.IsDir() {
                subSongs, err := ScanSongs(filePath)
                if err != nil {
                    continue
                }
                songs = append(songs, subSongs...)
            }
        }
    }
    return songs, err
}

func GetSongInfo(filePath string) (SongInfo, error) {
    if len(filePath) == 0 {
        return SongInfo{}, fmt.Errorf("Empty song file path was passed.\n")
    }
    s := SongInfo{ }
    s.FilePath = filePath
    file, err := os.Open(s.FilePath)
    if err != nil {
        return SongInfo{}, err
    }
    defer file.Close()
    id3Hdr := make([]byte, 6)
    n, err := file.Read(id3Hdr)
    if err != nil || n < len(id3Hdr) {
        return SongInfo{}, fmt.Errorf("Invalid ID3 header.\n")
    }
    if id3Hdr[3] != 3 {
        return SongInfo{}, fmt.Errorf("Unsupported ID3 header.\n")
    }
    hdrLen := make([]byte, 4)
    n, err = file.Read(hdrLen)
    if err != nil {
        return SongInfo{}, err
    }
    hdrSize := (int(hdrLen[0]) << 21) |
               (int(hdrLen[1]) << 14) |
               (int(hdrLen[2]) <<  7) |
               int(hdrLen[3])
    hdrData := make([]byte, hdrSize)
    _, err = file.Read(hdrData)
    if err != nil {
        return SongInfo{}, err
    }
    readFields := 1
    for h := 0; (h + 4) < len(hdrData) && readFields < kSongInfoFieldsNr; h++ {
        needle := string(hdrData[h:])
        var field *string = nil
        if len(s.Title) == 0 && strings.HasPrefix(needle, "TIT2") {
            field = &s.Title
        } else if len(s.Album) == 0 && strings.HasPrefix(needle, "TALB") {
            field = &s.Album
        } else if len(s.Artist) == 0 && strings.HasPrefix(needle, "TPE") {
            field = &s.Artist
        } else if len(s.TrackNumber) == 0 && strings.HasPrefix(needle, "TRCK") {
            field = &s.TrackNumber
        } else if len(s.Genre) == 0 && strings.HasPrefix(needle, "TCON") {
            field = &s.Genre
        } else if len(s.Year) == 0 && strings.HasPrefix(needle, "TYER") {
            field = &s.Year
        } else if len(s.AlbumCover) == 0 && strings.HasPrefix(needle, "APIC") {
            field = &s.AlbumCover
        }
        if field != nil {
            readFields++
            needleSize := (int(needle[4]) << 24) |
                          (int(needle[5]) << 16) |
                          (int(needle[6]) <<  8) |
                          int(needle[7]) - 1
            h += needleSize
            if needle[10] == 0 {
                *field = string(needle[11:])[:needleSize]
                h += 1
                *field = strings.Trim(*field, "\x00 ")
            } else if (needle[11] == 0xFF && needle[12] == 0xFE) || (needle[11] == 0xFE && needle[12] == 0xFF) {
                *field = string(needle[13:])[:needleSize - 2]
                if field != &s.AlbumCover {
                    *field = utfToAscii(*field)
                }
                h += 3
            }
            if field == &s.TrackNumber {
                subReg := regexp.MustCompile("/.*")
                s.TrackNumber = subReg.ReplaceAllString(s.TrackNumber, "")
            } 
        }
    }
    if len(s.Artist) == 0 {
        s.Artist = "[Unknown Artist]"
    }
    if len(s.Album) == 0 {
        s.Album = "[Unknown Album]"
    }
    if len(s.Title) == 0 {
        s.Title = "[Unknown Track]"
    }
    if len(s.TrackNumber) == 0 {
        s.TrackNumber = "0"
    }
    return s, nil
}

func utfToAscii(utfStr string) string {
    utfStrLen := len(utfStr)
    if utfStrLen == 0 {
        return ""
    }
    mbStr := make([]byte, utf8.RuneCountInString(utfStr) >> 1)
    m := 0
    u := 0
    for u < utfStrLen && utfStr[u] != 0x00 {
        mbStr[m] = byte(utfStr[u])
        m += 1
        u += 2
    }
    for m = 0; m < len(mbStr); m++ {
        if mbStr[m] == 0x00 {
            break
        }
    }
    return strings.Trim(string(mbStr), "\x00 ")
}

func stripNullEndings(str string) string {
    for s := 0; s < len(str); s++ {
        if str[s] == 0 {
            return str[0:s]
        }
    }
    return str
}