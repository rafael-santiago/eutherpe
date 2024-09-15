//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package mplayer

import (
    "os"
    "os/exec"
    "fmt"
    "strings"
    "regexp"
    "path"
    "crypto/sha256"
)

// QUESTION(Rafael): Shall we fix a maximum cover size?
//const (
//    kMaximumCoverSize = 1<<20
//)

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

func ConvertSongs(basePath string, customPath... string) error {
    files, err := os.ReadDir(basePath)
    if err != nil {
        return err
    }
    for _, file := range files {
        fileName := file.Name()
        filePath := path.Join(basePath, fileName)
        ext := path.Ext(strings.ToLower(fileName))
        if ext == ".m4a" || ext == ".mp4" {
            err = ConvertToMP3(filePath, customPath...)
            if err != nil {
                fmt.Fprintf(os.Stderr, "warn: error while converting '%s'.\n", filePath)
            }
        } else {
            stat, err := os.Stat(filePath)
            if err == nil && stat.IsDir() {
                err = ConvertSongs(filePath, customPath...)
            }
        }
    }
    return nil
}

func ScanSongs(basePath string, coversCacheRootPath ...string) ([]SongInfo, error) {
    files, err := os.ReadDir(basePath)
    if err != nil {
        return make([]SongInfo, 0), err
    }
    songs := make([]SongInfo, 0)
    for _, file := range files {
        fileName := strings.ToLower(file.Name())
        if path.Ext(fileName) == ".mp3" {
            songInfo, err := GetSongInfo(path.Join(basePath, file.Name()), coversCacheRootPath...)
            if err != nil {
                continue
            }
            songs = append(songs, songInfo)
        } else {
            filePath := path.Join(basePath, fileName)
            stat, err := os.Stat(filePath)
            if err == nil && stat.IsDir() {
                subSongs, err := ScanSongs(filePath, coversCacheRootPath...)
                if err != nil {
                    continue
                }
                songs = append(songs, subSongs...)
            }
        }
    }
    if err == nil && len(coversCacheRootPath) > 0 {
        cleanUpUnusedCovers(coversCacheRootPath[0], songs)
    }
    return songs, err
}

func GetSongInfo(filePath string, coversCacheRootPath ...string) (SongInfo, error) {
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
    id3Hdr := make([]byte, 12)
    n, err := file.Read(id3Hdr)
    if n < 12 {
        err = fmt.Errorf("No enough bytes.")
    }
    if err != nil {
        return  SongInfo{}, err
    }
    if string(id3Hdr[4:]) == "ftypM4A " ||
       string(id3Hdr[4:]) == "ftypisom" ||
       string(id3Hdr[4:]) == "ftypiso2" {
        return getSongInfoFromM4A(filePath, coversCacheRootPath...)
    }
    if id3Hdr[0] == 0xFF && id3Hdr[1] == 0xFB {
        return getSongInfoFromIDv1(filePath, coversCacheRootPath...)
    }
    if err != nil || n < len(id3Hdr) || (id3Hdr[0] != 'I' && id3Hdr[1] != 'D' && id3Hdr[2] != '3') {
        return SongInfo{}, fmt.Errorf("Invalid ID3 header.")
    }
    id3V := id3Hdr[3]
    if id3V != 2 && id3V != 3 && id3V != 4 {
        return SongInfo{}, fmt.Errorf("Unsupported ID3 header.")
    }
    _, err = file.Seek(6, 0)
    if err != nil {
        return SongInfo{}, err
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
        "TT2",
        "TALB",
        "TAL",
        "TPE1",
        "TPE2",
        "TP1",
        "TRCK",
        "TRK",
        "TCON",
        "TCO",
        "TYER",
        "TYE",
        "APIC",
        "PIC",
    }
    strHdrData := string(hdrData)
    shouldUseCachedCovers := (len(coversCacheRootPath) > 0 && len(coversCacheRootPath[0]) > 0)
    for _, info := range kWantedInfo {
        h := strings.Index(strHdrData, info)
        if h == -1 {
            continue
        }
        var field *string = nil
        switch info {
            case "TIT2", "TT2":
                field = &s.Title
                break
            case "TALB", "TAL":
                field = &s.Album
                break
            case "TPE1", "TPE2", "TP1":
                field = &s.Artist
                break
            case "TRCK", "TRK":
                field = &s.TrackNumber
                break
            case "TCON", "TCO":
                field = &s.Genre
                break
            case "TYER", "TYE":
                field = &s.Year
                break
            case "APIC", "PIC":
                field = &s.AlbumCover
                break
            default:
                continue
        }
        if field != nil && len(*field) == 0 {
            needle := string(hdrData[h:])
            // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            // WARN(Rafael): Some MP3 that I messed with was ID3v2.4.0 (id3V == 4) BUT           !!
            //               the APIC field was being recorded as ID3v2.3.0 (id3V should be 3!). !!
            //               The most dumb folks in the UNIVERSE had designed MP3 ID3, what      !!
            //               a mess!! How people can be so DUMB and UNFIT, Gosh!!!!!!!           !!
            //               What a damn HELL! Go learn how to think before writing code and     !!
            //               most critical: before creating standards! If you are unfit on       !!
            //               thinking and/or designing, just follow, do not try to lead          !!
            //               anyone or anything, please. It will be the greateast contribution   !!
            //               you could do for the whole f_cking World! More than a lousy shitty  !!
            //               standard, that endlessly haunt folks by propagating your dumbness   !!
            //               and lousy naive decisions everywhere, really! Thank U and farewell! !!
            // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            var needleSize int
            //var startOff int
            switch id3V {
                case 4:
                    needleSize = (int(needle[4]) << 21) |
                                 (int(needle[5]) << 14) |
                                 (int(needle[6]) <<  7) |
                                 int(needle[7])
                    //startOff = 11
                    break
                case 3:
                    needleSize = (int(needle[4]) << 24) |
                                 (int(needle[5]) << 16) |
                                 (int(needle[6]) <<  8) |
                                 int(needle[7])
                    //startOff = 13
                    break
                case 2:
                    needleSize = (int(needle[3]) << 16) |
                                 (int(needle[4]) <<  8) |
                                 int(needle[5])
                    break
            }
            needleSize -= 1
            //fmt.Println("info: ", info)
            //fmt.Println("needleSize = ", needleSize)
            if needleSize > len(needle) {
                // INFO(Rafael): Ahhh this is a mess, cross your fingers for it at least reproducing!
                continue
            }
            h += needleSize
            if field != &s.AlbumCover && (needle[10] == 0 && id3V != 2) || (needle[9] == 0 &&  id3V == 4 && needle[11] != 0xFF && needle[12] != 0xFE) {
                *field = string(needle[11:])[:needleSize]
                h += 1
                *field = strings.Trim(*field, "\x00 ")
            } else if field != &s.AlbumCover && id3V != 2 && (needle[11] == 0xFF && needle[12] == 0xFE) || (needle[11] == 0xFE && needle[12] == 0xFF) {
                *field = string(needle[13:])[:needleSize - 2]
                if field != &s.AlbumCover {
                    *field = utfToAscii(*field)
                }
                h += 3
            } else if field == &s.AlbumCover {
                if mimeTypeIndex := strings.Index(needle[:20], "image/"); mimeTypeIndex > -1 {
                    blobData := needle[mimeTypeIndex + 6:]
                    startOff := -1
                    if strings.HasPrefix(blobData, "jpeg") {
                        startOff = strings.Index(blobData[:20], "\xFF\xD8\xFF\xE0")
                        if startOff == -1 {
                            startOff = strings.Index(blobData[:20], "\xFF\xD8\xFF\x00\xE0")
                        }
                    }
                    if startOff > -1 {
                        *field = blobData[startOff:]
                    } else {
                        // INFO(Rafael): We will still give it a try a little more further...
                        *field = needle[mimeTypeIndex:]
                    }
                } else if id3V == 2 {
                    *field = string(needle[7:])[:needleSize]
                }
            } else if id3V == 2 {
                *field = string(needle[7:])[:needleSize]
                if field != &s.AlbumCover {
                    nullByteIndex := strings.Index(*field, "\x00")
                    if nullByteIndex > -1 {
                        *field = (*field)[:nullByteIndex]
                    }
                }
            }
            if field == &s.TrackNumber {
                subReg := regexp.MustCompile("(^0|/.*)")
                s.TrackNumber = subReg.ReplaceAllString(s.TrackNumber, "")
            } else if field == &s.AlbumCover && len(s.AlbumCover) > 0 {
                a := 0
                mimeType := string(s.AlbumCover[0:20])
                isJPEG := strings.HasPrefix(mimeType, "image/jpeg") ||
                          strings.HasPrefix(mimeType, "image/jpg") ||
                          strings.HasPrefix(s.AlbumCover, "JPG")
                isPNG := !isJPEG && (strings.HasPrefix(mimeType, "image/png") ||
                                     strings.HasPrefix(s.AlbumCover, "PNG"))
                for ; isPNG && a < len(s.AlbumCover) && s.AlbumCover[a] != 0x89; a++ {
                }
                for skp := 0; isJPEG && skp < 2; skp++ {
                    for ; a < len(s.AlbumCover); a++ {
                        if s.AlbumCover[a] == 0x00 {
                            break
                        }
                    }
                    a++
                }
                if isJPEG && s.AlbumCover[0] != 0xFF && s.AlbumCover[1] != 0xD8 {
                    for ; a < len(s.AlbumCover); a++ {
                        if s.AlbumCover[a] == 0xFF {
                            break
                        }
                    }
                }
                if a > len(s.AlbumCover) {
                    s.AlbumCover = ""
                } else if !shouldUseCachedCovers {
                    s.AlbumCover = s.AlbumCover[a:]
                } else {
                    coverBlob := []byte(s.AlbumCover[a:])
                    var coverId string
                    if !isAlbumCoverCached(coverBlob, coversCacheRootPath[0], &coverId) {
                        makeAlbumCoverCache(coversCacheRootPath[0], coverId, coverBlob)
                    }
                    s.AlbumCover = "blob-id=" + coverId
                }
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
        s.TrackNumber = getTrackNumberFromFileName(filePath)
    }
    if len(s.AlbumCover) == 0 {
        s.AlbumCover = getAlbumCoverFromRootPath(filePath, coversCacheRootPath...)
    }
    normalizeSongInfo(&s)
    return s, nil
}

func getSongInfoFromIDv1(filePath string, coversCacheRootPath ...string) (SongInfo, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return SongInfo{}, err
    }
    if len(data) < 128 {
        return SongInfo{}, fmt.Errorf("Not enough bytes in supposed ID1 header.")
    }
    idv1Data := data[len(data) - 128:]
    if idv1Data[0] != 'T' && idv1Data[1] != 'A' && idv1Data[2] != 'G' {
        return SongInfo{}, fmt.Errorf("Invalid ID1 header.")
    }
    if len(idv1Data) < 97 {
        return SongInfo{}, fmt.Errorf("ID1 header seems corrupted.")
    }
    type IDv1MetaDataCtx struct {
        Title []byte
        Artist []byte
        Album []byte
        Year []byte
    }
    var IDv1MetaData IDv1MetaDataCtx
    IDv1MetaData.Title = make([]byte, 30)
    copy(IDv1MetaData.Title, idv1Data[3:])
    IDv1MetaData.Artist = make([]byte, 30)
    copy(IDv1MetaData.Artist, idv1Data[33:])
    IDv1MetaData.Album = make([]byte, 30)
    copy(IDv1MetaData.Album, idv1Data[63:])
    IDv1MetaData.Year = make([]byte, 4)
    copy(IDv1MetaData.Year, idv1Data[93:])
    songInfo := SongInfo{ Title: getStringFromNullTerminatedString(IDv1MetaData.Title),
                          Album: getStringFromNullTerminatedString(IDv1MetaData.Album),
                          Artist: getStringFromNullTerminatedString(IDv1MetaData.Artist),
                          Year: getStringFromNullTerminatedString(IDv1MetaData.Year),
                          FilePath: filePath,
                          TrackNumber: getTrackNumberFromFileName(filePath), }
    normalizeSongInfo(&songInfo)
    songInfo.AlbumCover = getAlbumCoverFromRootPath(filePath, coversCacheRootPath...)
    return songInfo, nil
}

func getSongInfoFromM4A(filePath string, coversCacheRootPath ...string) (SongInfo, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return SongInfo{}, err
    }
    shouldUseCachedCovers := (len(coversCacheRootPath) > 0 && len(coversCacheRootPath[0]) > 0)
    songInfo := SongInfo{}
    type ParserProgram struct {
        Tag string
        Data *string
    }
    songInfo.FilePath = filePath
    parserProgram := []ParserProgram {
        { "\xA9nam", &songInfo.Title },
        { "\xA9ART", &songInfo.Artist },
        { "\xA9alb", &songInfo.Album },
        { "trkn", &songInfo.TrackNumber },
        { "day", &songInfo.Year },
        { "covr", &songInfo.AlbumCover },
    }
    fileBuf := string(data)
    for _, step := range parserProgram {
        pos := strings.Index(fileBuf, step.Tag)
        if pos == -1 {
            continue
        }
        subBuf := fileBuf[pos:]
        dataBuf := subBuf[len(step.Tag):]
        dataSize := ((int(dataBuf[0]) << 24) |
                     (int(dataBuf[1]) << 16) |
                     (int(dataBuf[2]) <<  8) |
                     int(dataBuf[3]))
        if dataSize > len(dataBuf) {
            continue
        }
        infoBuf := dataBuf[4:dataSize]
        if !strings.HasPrefix(infoBuf, "data") {
            continue
        }
        dataSize = len(infoBuf)
        var d int
        if step.Tag != "trkn" && step.Tag != "covr" {
            d = dataSize - 1
            for ; d >= 0 && infoBuf[d] != 0x00; d-- {
            }
            if d < 0 || (d + 1) >= dataSize {
                continue
            }
            *step.Data = infoBuf[d+1:]
        } else if step.Tag == "trkn" {
            infoBuf = infoBuf[4:]
            for d = 0 ; d < dataSize && infoBuf[d] == 0; d++ {
            }
            *step.Data = fmt.Sprintf("%d", infoBuf[d])
        } else if step.Tag == "covr" {
            infoBuf = infoBuf[4:]
            pngPos := strings.Index(infoBuf, "PNG")
            if pngPos == -1 {
                continue
            }
            coverBlob := []byte(infoBuf[pngPos - 1:])
            //if (len(coverBlob) > kMaximumCoverSize) {
            //    coverBlob = nil
            //    continue
            //}
            if !shouldUseCachedCovers {
                *step.Data = string(coverBlob)
            } else {
                var coverId string
                if !isAlbumCoverCached(coverBlob, coversCacheRootPath[0], &coverId) {
                    makeAlbumCoverCache(coversCacheRootPath[0], coverId, coverBlob)
                }
                *step.Data = "blob-id=" + coverId
            }
        }
    }
    if len(songInfo.Artist) == 0 {
        songInfo.Artist = "[Unknown Artist]"
    }
    if len(songInfo.Album) == 0 {
        songInfo.Album = "[Unknown Album]"
    }
    if len(songInfo.Title) == 0 {
        songInfo.Title = path.Base(filePath)
    }
    if len(songInfo.TrackNumber) == 0 {
        songInfo.TrackNumber = getTrackNumberFromFileName(filePath)
    }
    if len(songInfo.AlbumCover) == 0 {
        songInfo.AlbumCover = getAlbumCoverFromRootPath(filePath, coversCacheRootPath...)
    }
    normalizeSongInfo(&songInfo)
    return songInfo, nil
}

func normalizeStr(str string) string {
    var normStr string
    for s, annoyingRune := range str {
        //fmt.Printf("%d - %v\n", int(str[s]), annoyingRune)
        switch annoyingRune {
            case 226, 224, 225, 227, 228, 229:
                normStr += "a"
                break
            case 194, 192, 193, 195, 196, 197:
                normStr += "A"
                break
            case 242, 243, 244, 245, 246, 210, 211, 212, 213, 214:
                normStr += "O"
                break
            case 232, 233, 235, 234:
                normStr += "e"
                break
            case 200, 201, 203, 202:
                normStr += "E"
                break
            case 236, 237, 239:
                normStr += "i"
                break
            case 204, 205, 207:
                normStr += "I"
                break
            case 231:
                normStr += "c"
                break
            case 199:
                normStr += "C"
                break
            case 152:
                normStr += "y"
                break
            case 241:
                normStr += "n"
                break
            case 209:
                normStr += "N"
                break
            case 176:
                normStr += "c"
                break
            case 39, 8217:
                normStr += "_"
                break
            case 249, 250, 252:
                normStr += "u"
                break
            case 217, 218, 220:
                normStr += "U"
                break

            default:
                // TODO(Rafael): Find a way of improve it on, it is totally clumsy.
                //               BUT it works!
                switch str[s] {
                    case 226, 224, 225, 227, 228, 229:
                        normStr += "a"
                        break
                    case 194, 192, 193, 195, 196, 197:
                        normStr += "A"
                        break
                    case 242, 243, 244, 245, 246, 210, 211, 212, 213, 214:
                        normStr += "O"
                        break
                    case 232, 233, 235, 234:
                        normStr += "e"
                        break
                    case 200, 201, 203, 202:
                        normStr += "E"
                        break
                    case 236, 237, 239:
                        normStr += "i"
                        break
                    case 204, 205, 207:
                        normStr += "I"
                        break
                    case 231:
                        normStr += "c"
                        break
                    case 199:
                        normStr += "C"
                        break
                    case 152:
                        normStr += "y"
                        break
                    case 241:
                        normStr += "n"
                        break
                    case 209:
                        normStr += "N"
                        break
                    case 176:
                        normStr += "c"
                        break
                    case 39:
                        normStr += "_"
                        break
                    case 249, 250, 252:
                        normStr += "u"
                        break
                    case 217, 218, 220:
                        normStr += "U"
                        break
                    case 167:
                        continue
                    default:
                        normStr += string(annoyingRune)
                        break
                }
                break
        }
    }
    return normStr
}

func normalizeSongInfo(songInfo *SongInfo) {
    songInfo.Artist = strings.Replace(songInfo.Artist, "/", "-", -1)
    songInfo.Artist = strings.Replace(songInfo.Artist, "\"", "'", -1)
    songInfo.Album = strings.Replace(songInfo.Album, "/", "-", -1)
    songInfo.Album = strings.Replace(songInfo.Album, "\"", "'", -1)
    songInfo.Title = strings.Replace(songInfo.Title, "/", "-", -1)
    songInfo.Title = strings.Replace(songInfo.Title, "\"", "'", -1)
    songInfo.Artist = strings.ToLower(normalizeStr(songInfo.Artist))
    songInfo.Album = strings.ToLower(normalizeStr(songInfo.Album))
    songInfo.Title = strings.ToLower(normalizeStr(songInfo.Title))
}

func getTrackNumberFromFileName(filePath string) string {
    var startOff int = -1
    var endOff int = -1
    fileName := path.Base(filePath)
    for f, char := range fileName {
        if char >= '0' && char <= '9' {
            if startOff == -1 {
                startOff = f
            }
        } else if startOff > -1 {
            endOff = f
            break
        }
    }
    if startOff == -1 || endOff == -1 {
        return "0"
    }
    return string(fileName[startOff:endOff])
}

func getStringFromNullTerminatedString(cStr []byte) string {
    var off int = 0
    for ; off < len(cStr) && cStr[off] != 0x00; off++ {
    }
    return string(cStr[:off])
}

func utfToAscii(utfStr string) string {
    utfStrLen := len(utfStr)
    if utfStrLen == 0 {
        return ""
    }
    u := 0
    var mbStr string
    for  u  < utfStrLen && utfStr[u] != 0x00 {
        mbStr += string(utfStr[u])
        u += 2
    }
    return strings.Trim(mbStr, "\x00 ")
}

func getAlbumCoverId(blob []byte) string {
    imageHash :=  sha256.Sum256(blob)
    return fmt.Sprintf("%x", imageHash)
}

func isAlbumCoverCached(blob []byte, coversCacheRootPath string, imageHash *string) bool {
    (*imageHash) = getAlbumCoverId(blob)
    if len(*imageHash) == 0 {
        return false
    }
    err := os.MkdirAll(coversCacheRootPath, 0777)
    if err != nil {
        return false
    }
    _, err = os.Stat(path.Join(coversCacheRootPath, *imageHash))
    return (err == nil)
}

func resizeAlbumCover(sourcePath, destPath string, width, height uint) error {
    fmt.Println(sourcePath, destPath)
    return exec.Command("convert", sourcePath, "-resize",
                        fmt.Sprintf("%dx%d", width, height),
                        destPath).Run()
}

func makeAlbumCoverCache(coversCacheRootPath, imageHash string, blob []byte) {
    tempCover, err := os.CreateTemp("", imageHash)
    if err != nil {
        return
    }
    tempCoverFilePath := tempCover.Name()
    defer os.Remove(tempCover.Name())
    defer tempCover.Close()
    tempCover.Write(blob)
    if resizeAlbumCover(tempCoverFilePath,
                        path.Join(coversCacheRootPath, imageHash),
                        500, 500) != nil {
        fmt.Fprintf(os.Stderr, "warn: Unable to resize album cover image. Using the original size.\n")
        os.WriteFile(path.Join(coversCacheRootPath, imageHash), blob, 0777)
    }
}

func cleanUpUnusedCovers(coversCacheRootPath string, songs []SongInfo) {
    coversListing, err := os.ReadDir(path.Join(coversCacheRootPath))
    if err != nil || len(coversListing) < 100 {
        return
    }
    covers := make(map[string]bool)
    for _, coverFile := range coversListing {
        for _, song := range songs {
            if strings.HasPrefix(song.AlbumCover, "blob-id=") && song.AlbumCover[8:] == coverFile.Name() {
                covers[coverFile.Name()] = true
                break
            }
        }
    }
    for coverId, used := range covers {
        if !used {
            os.Remove(path.Join(coversCacheRootPath, coverId))
        }
    }
}

func getAlbumCoverFromRootPath(filePath string, coversCacheRootPath ...string) string {
    dirPath := path.Dir(filePath)
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return ""
    }
    coverBlob := make([]byte, 0)
    for _, file := range files {
        fileName := strings.ToLower(file.Name())
        if path.Ext(fileName) == ".jpg" ||
           path.Ext(fileName) == ".jpeg" ||
           path.Ext(fileName) == ".png" {
            coverBlob, err = os.ReadFile(path.Join(dirPath, file.Name()))
            if err != nil {
                break
            }
        }
    }
    shouldUseCachedCovers := (len(coversCacheRootPath) > 0 && len(coversCacheRootPath[0]) > 0)
    if !shouldUseCachedCovers {
        return string(coverBlob)
    }
    var coverId string
    if !isAlbumCoverCached(coverBlob, coversCacheRootPath[0], &coverId) {
        makeAlbumCoverCache(coversCacheRootPath[0], coverId, coverBlob)
    }
    return ("blob-id=" + coverId)
}
