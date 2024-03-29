package mplayer

import (
    "os"
    "fmt"
    "strings"
    "regexp"
    "path"
    "unicode/utf8"
    "crypto/sha256"
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

func ScanSongs(basePath string, coversCacheRootPath ...string) ([]SongInfo, error) {
    files, err := os.ReadDir(basePath)
    if err != nil {
        return make([]SongInfo, 0), err
    }
    songs := make([]SongInfo, 0)
    for _, file := range files {
        fileName := file.Name()
        if path.Ext(fileName) == ".mp3" || path.Ext(fileName) == ".mp4" {
            songInfo, err := GetSongInfo(path.Join(basePath, fileName), coversCacheRootPath...)
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
    id3Hdr := make([]byte, 6)
    n, err := file.Read(id3Hdr)
    if id3Hdr[0] == 0xFF && id3Hdr[1] == 0xFB {
        return getSongInfoFromIDv1(filePath)
    }
    if err != nil || n < len(id3Hdr) || (id3Hdr[0] != 'I' && id3Hdr[1] != 'D' && id3Hdr[2] != '3') {
        return SongInfo{}, fmt.Errorf("Invalid ID3 header.")
    }
    id3V := id3Hdr[3]
    if id3V != 3 && id3V != 4 {
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
    shouldUseCachedCovers := (len(coversCacheRootPath) > 0 && len(coversCacheRootPath[0]) > 0)
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
            if id3V == 4 {
                needleSize = (int(needle[4]) << 21) |
                             (int(needle[5]) << 14) |
                             (int(needle[6]) <<  7) |
                             int(needle[7])
            } else {
                needleSize = (int(needle[4]) << 24) |
                             (int(needle[5]) << 16) |
                             (int(needle[6]) <<  8) |
                             int(needle[7])
            }
            needleSize -= 1
            if needleSize > len(needle) {
                // INFO(Rafael): Ahhh this is a mess, cross your fingers for it at least reproducing!
                continue
            }
            h += needleSize
            if needle[10] == 0 || (needle[9] == 0 &&  id3V == 4 && needle[11] != 0xFF && needle[12] != 0xFE) {
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
            } else if field == &s.AlbumCover && len(s.AlbumCover) > 0 {
                a := 0
                mimeType := string(s.AlbumCover[0:20])
                isJPEG := strings.HasPrefix(mimeType, "image/jpeg")
                isPNG := !isJPEG && strings.HasPrefix(mimeType, "image/png")
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
    return s, nil
}

func getSongInfoFromIDv1(filePath string) (SongInfo, error) {
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
    return SongInfo{ Title: getStringFromNullTerminatedString(IDv1MetaData.Title),
                     Album: getStringFromNullTerminatedString(IDv1MetaData.Album),
                     Artist: getStringFromNullTerminatedString(IDv1MetaData.Artist),
                     Year: getStringFromNullTerminatedString(IDv1MetaData.Year),
                     FilePath: filePath,
                     TrackNumber: getTrackNumberFromFileName(filePath), }, nil
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

func makeAlbumCoverCache(coversCacheRootPath, imageHash string, blob []byte) {
    os.WriteFile(path.Join(coversCacheRootPath, imageHash), blob, 0777)
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