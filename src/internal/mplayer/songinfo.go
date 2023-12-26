package mplayer

import (
    "os"
    "fmt"
    "strings"
    "regexp"
    "path"
    "unicode/utf8"
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
        return SongInfo{}, fmt.Errorf("Empty song file path was passed.")
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
    if err != nil || n < len(id3Hdr) || (id3Hdr[0] != 'I' && id3Hdr[1] != 'D' && id3Hdr[2] != '3') {
        return SongInfo{}, fmt.Errorf("Invalid ID3 header.")
    }
    if id3Hdr[3] != 3 {
        return SongInfo{}, fmt.Errorf("Unsupported ID3 header.")
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
    kWantedInfo := []string {
        "TIT2",
        "TALB",
        "TPE",
        "TRCK",
        "TCON",
        "TYER",
        "APIC",
    }
    strHdrData := string(hdrData)
    for _, info := range kWantedInfo {
        h := strings.Index(strHdrData, info)
        if h == -1 {
            continue
        }
        var field *string = nil
        switch info {
            case "TIT2":
                field = &s.Title
                break
            case "TALB":
                field = &s.Album
                break
            case "TPE":
                field = &s.Artist
                break
            case "TRCK":
                field = &s.TrackNumber
                break
            case "TCON":
                field = &s.Genre
                break
            case "TYER":
                field = &s.Year
                break
            case "APIC":
                field = &s.AlbumCover
                break
            default:
                continue
        }
        if field != nil {
            needle := string(hdrData[h:])
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
                subReg := regexp.MustCompile("(^0|/.*)")
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
