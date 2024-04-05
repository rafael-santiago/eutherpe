package vars

import (
    "internal/mplayer"
    "internal/dj"
    "internal/bluebraces"
    "internal/storage"
    "internal/auth"
    "sync"
    "os/exec"
    "encoding/json"
    "encoding/base64"
    "os"
    "path"
    "fmt"
    "strings"
)

type EutherpeVars struct {
    APPName string
    ConfHome string
    HTTPd struct {
        Authenticated bool
        TLS bool
        AuthWatchdog *auth.AuthWatchdog
        HashKey string
        URLSchema string
        Addr string
        PubRoot string
        PubFiles []string
        IndexHTML string
        ErrorHTML string
        LoginHTML string
    }
    BluetoothDevices []bluebraces.BluetoothDevice
    StorageDevices []string
    CachedDevices struct {
        BlueDevId string
        MusicDevId string
    }
    Collection mplayer.MusicCollection
    Playlists []dj.Playlist
    Tags dj.Tags
    RenderedPlaylist string
    Player struct {
        NowPlaying mplayer.SongInfo
        UpNext []mplayer.SongInfo
        UpNextCurrentOffset int
        Handle *exec.Cmd
        UpNextBkp []mplayer.SongInfo
        Shuffle bool
        RepeatAll bool
        RepeatOne bool
        Stopped bool
        VolumeLevel uint
    }
    LastError error
    CurrentConfig string
    LastCommonTags []string
    LastSelection string
    mtx sync.Mutex
}

type eutherpeVarsCacheCtx struct {
    UpNext []mplayer.SongInfo
    Shuffle bool
    VolumeLevel uint
    RepeatOne bool
    RepeatAll bool
    BlueDevId string
    MusicDevId string
    UpNextCurrentOffset int
    Authenticated bool
    HashKey string
    TLS bool
}

func (e *EutherpeVars) Lock() {
    e.mtx.Lock()
}

func (e *EutherpeVars) Unlock() {
    e.mtx.Unlock()
}

func (e *EutherpeVars) toJSON() string {
    cachedData := eutherpeVarsCacheCtx { e.Player.UpNext,
                                         e.Player.Shuffle,
                                         e.Player.VolumeLevel,
                                         e.Player.RepeatOne,
                                         e.Player.RepeatAll,
                                         e.CachedDevices.BlueDevId,
                                         e.CachedDevices.MusicDevId,
                                         e.Player.UpNextCurrentOffset,
                                         e.HTTPd.Authenticated,
                                         e.HTTPd.HashKey,
                                         e.HTTPd.TLS, }
    if e.Player.Shuffle {
        cachedData.UpNext = e.Player.UpNextBkp
    }
    for u, _ := range cachedData.UpNext {
        isCachedAlbumCover := strings.HasPrefix(cachedData.UpNext[u].AlbumCover, "blob-id=")
        if !isCachedAlbumCover && len(cachedData.UpNext[u].AlbumCover) > 0 {
            cachedData.UpNext[u].AlbumCover = base64.StdEncoding.EncodeToString([]byte(cachedData.UpNext[u].AlbumCover))
        }
    }
    data, err := json.Marshal(&cachedData)
    if err != nil {
        return ""
    }
    return string(data)
}

func (e *EutherpeVars) fromJSON(filePath string) error {
    jsonData, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }
    var cachedData eutherpeVarsCacheCtx
    err = json.Unmarshal([]byte(jsonData), &cachedData)
    if err != nil {
        return err
    }
    for u, _ := range cachedData.UpNext {
        isCachedAlbumCover := strings.HasPrefix(cachedData.UpNext[u].AlbumCover, "blob-id=")
        if !isCachedAlbumCover && len(cachedData.UpNext[u].AlbumCover) > 0 {
            blob, _ := base64.StdEncoding.DecodeString(cachedData.UpNext[u].AlbumCover)
            cachedData.UpNext[u].AlbumCover = string(blob)
        }
    }
    e.Player.UpNext = cachedData.UpNext
    e.Player.UpNextBkp = cachedData.UpNext
    e.Player.Shuffle = cachedData.Shuffle
    e.Player.VolumeLevel = cachedData.VolumeLevel
    e.Player.RepeatOne = cachedData.RepeatOne
    e.Player.RepeatAll = cachedData.RepeatAll
    e.Player.UpNextCurrentOffset = cachedData.UpNextCurrentOffset
    e.CachedDevices.BlueDevId = cachedData.BlueDevId
    e.CachedDevices.MusicDevId = cachedData.MusicDevId
    e.HTTPd.Authenticated = cachedData.Authenticated
    e.HTTPd.HashKey = cachedData.HashKey
    if len(e.HTTPd.HashKey) == 0 {
        e.HTTPd.HashKey = auth.HashKey("music")
    }
    e.HTTPd.TLS = cachedData.TLS
    if cachedData.TLS {
        e.HTTPd.URLSchema = "https"
    } else {
        e.HTTPd.URLSchema = "http"
    }
    return nil
}

func (e *EutherpeVars) SaveSession() error {
    playerSettings := e.toJSON()
    if len(playerSettings) == 0 {
        return fmt.Errorf("Unable to serialize player settings.")
    }
    err := os.WriteFile(path.Join(e.ConfHome, EutherpePlayerCache), []byte(playerSettings), 0666)
    if err != nil {
        return nil
    }
    err = e.SaveCollection()
    if err != nil {
        return err
    }
    err = e.SavePlaylists()
    if err != nil {
        return err
    }
    return e.SaveTags()
}


func (e *EutherpeVars) RestoreSession() error {
    err := e.fromJSON(path.Join(e.ConfHome, EutherpePlayerCache))
    if err != nil {
        return err
    }
    err = e.LoadCollection()
    if err != nil {
        return err
    }
    return nil
}

func (e *EutherpeVars) SaveCollection() error {
    if len(e.Collection) == 0 || len(e.CachedDevices.MusicDevId) == 0 {
        return nil
    }
    cacheFilePath := path.Join(e.ConfHome, EutherpeLastCollectionsHome)
    err := os.MkdirAll(cacheFilePath, 0777)
    if err != nil {
        fmt.Println(err)
        return err
    }
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(e.CachedDevices.MusicDevId)
    cacheFilePath = path.Join(cacheFilePath, musicDevSerial)
    return os.WriteFile(cacheFilePath, []byte(e.Collection.ToJSON()), 0777)
}

func (e *EutherpeVars) LoadCollection() error {
    if len(e.CachedDevices.MusicDevId) == 0 {
        return nil
    }
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(e.CachedDevices.MusicDevId)
    cacheFilePath := path.Join(e.ConfHome, EutherpeLastCollectionsHome, musicDevSerial)
    _, err := os.Stat(cacheFilePath)
    if err != nil {
        return nil
    }
    err = e.Collection.FromJSON(cacheFilePath)
    if err != nil {
        return err
    }
    e.Playlists = make([]dj.Playlist, 0)
    playlistsRootPath := path.Join(e.ConfHome, EutherpePlaylistsHome, musicDevSerial)
    _, err = os.Stat(playlistsRootPath)
    if err != nil {
        return nil
    }
    playlistsFiles, err := os.ReadDir(playlistsRootPath)
    for _, playlistFile := range playlistsFiles {
        playlist := dj.Playlist{}
        err := playlist.LoadFrom(path.Join(playlistsRootPath, playlistFile.Name()))
        if err != nil {
            return err
        }
        e.Playlists = append(e.Playlists, playlist)
    }
    return e.LoadTags()
}

func (e *EutherpeVars) SavePlaylists() error {
    if len(e.Playlists) == 0 {
        return nil
    }
    for p, _ := range e.Playlists {
        err := e.SavePlaylist(&e.Playlists[p])
        if err != nil {
            return err
        }
    }
    return nil
}

func (e *EutherpeVars) SavePlaylist(playlist *dj.Playlist) error {
    if playlist == nil {
        return nil
    }
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(e.CachedDevices.MusicDevId)
    playlistsRootPath := path.Join(e.ConfHome, EutherpePlaylistsHome, musicDevSerial)
    err := os.MkdirAll(playlistsRootPath, 0777)
    if err != nil {
        return err
    }
    return playlist.SaveTo(path.Join(playlistsRootPath, playlist.Name))
}

func (e *EutherpeVars) RemovePlaylistFromDisk(playlistName string) error {
    if len(playlistName) == 0 {
        return fmt.Errorf("No playlist name was provided.")
    }
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(e.CachedDevices.MusicDevId)
    playlistsRootPath := path.Join(e.ConfHome, EutherpePlaylistsHome, musicDevSerial)
    return os.Remove(path.Join(playlistsRootPath, playlistName))
}

func (e *EutherpeVars) SaveTags() error {
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(e.CachedDevices.MusicDevId)
    tagsRootPath := path.Join(e.ConfHome, EutherpeTagsHome)
    err := os.MkdirAll(tagsRootPath, 0777)
    if err != nil {
        return err
    }
    return e.Tags.SaveTo(path.Join(tagsRootPath, musicDevSerial))
}

func (e *EutherpeVars) LoadTags() error {
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(e.CachedDevices.MusicDevId)
    deviceTagsFilePath := path.Join(e.ConfHome, EutherpeTagsHome, musicDevSerial)
    _, err := os.Stat(deviceTagsFilePath)
    if err != nil {
        return nil
    }
    return e.Tags.LoadFrom(deviceTagsFilePath)
}

func (e *EutherpeVars) GetCoversCacheRootPath() string {
    return path.Join(e.ConfHome, EutherpeCoversHome)
}

const EutherpeActionId = "action"

// INFO(Rafael): Actions from "Music" sheet.

const EutherpeMusicRemoveId = "music-remove"
const EutherpeMusicMoveUpId = "music-moveup"
const EutherpeMusicMoveDownId = "music-movedown"
const EutherpeMusicClearAllId = "music-clearall"
const EutherpeMusicShuffleId = "music-shuffle"
const EutherpeMusicRepeatAllId = "music-repeatall"
const EutherpeMusicRepeatOneId = "music-repeatone"
const EutherpeMusicPlayId = "music-play"
const EutherpeMusicStopId = "music-stop"
const EutherpeMusicNextId = "music-next"
const EutherpeMusicLastId = "music-last"
const EutherpeMusicSetVolumeId = "music-setvolume"

// INFO(Rafael): Actions from "Collection" sheet.

const EutherpeCollectionAddSelectionToNextId = "collection-addselectiontonext"
const EutherpeCollectionAddSelectionToUpNextId = "collection-addselectiontoupnext"
const EutherpeCollectionAddSelectionToPlaylistId = "collection-addselectiontoplaylist"
const EutherpeCollectionTagSelectionAsId = "collection-tagselectionas"
const EutherpeCollectionUntagSelectionsId = "collection-untagselections"
const EutherpeCollectionPlayByGivenTagsId = "collection-playbygiventags"

// INFO(Rafael): Actions from "Playlists" sheet.

const EutherpePlaylistCreateId = "playlist-create"
const EutherpePlaylistRemoveId = "playlist-remove"
const EutherpePlaylistShowId = "playlist-show"
const EutherpePlaylistMoveUpId = "playlist-moveup"
const EutherpePlaylistMoveDownId = "playlist-movedown"
const EutherpePlaylistClearAllId = "playlist-clearall"
const EutherpePlaylistRemoveSongsId = "playlist-removesongs"
const EutherpePlaylistReproduceId = "playlist-reproduce"
const EutherpePlaylistReproduceSelectedOnesId = "playlist-reproduceselectedones"

// INFO(Rafael): Actions from "Storage" sheet.

const EutherpeStorageListId = "storage-list"
const EutherpeStorageScanId = "storage-scan"
const EutherpeStorageSetId = "storage-set"

// INFO(Rafael): Actions from "Bluetooth" sheet.

const EutherpeBluetoothProbeDevicesId = "bluetooth-probedevices"
const EutherpeBluetoothPairId = "bluetooth-pair"
const EutherpeBluetoothUnpairId = "bluetooth-unpair"
const EutherpeBluetoothTrustId = "bluetooth-trust"
const EutherpeBluetoothUntrustId = "bluetooth-untrust"

// INFO(Rafael): Actions from "Settings" sheet.

const EutherpeSettingsFlickAuthModeId = "settings-flickauthmode"
const EutherpeSettingsChangePassphraseId = "settings-changepassphrase"
const EutherpeSettingsFlickHTTPSModeId = "settings-flickhttpsmode"
const EutherpeSettingsGenerateCertificateId = "settings-generatecertificate"

const EutherpePlayerStatusId = "player-status"
const EutherpeGetCommonTagsId = "get-commontags"

const EutherpeAuthenticateId = "authenticate"

const EutherpePostFieldSelection = "selection"
const EutherpePostFieldPlaylist = "playlist"
const EutherpePostFieldStorageDevice = "storage-device"
const EutherpePostFieldBluetoothDevice = "bluetooth-device"
const EutherpePostFieldVolumeLevel = "volume-level"
const EutherpePostFieldLastError = "last-error"
const EutherpePostFieldTags = "tags"
const EutherpePostFieldAmount = "amount"
const EutherpePostFieldRemoteAddr = "remote-addr"
const EutherpePostFieldPassword = "password"
const EutherpePostFieldNewPassword = "new-password"

// INFO(Rafael): Template markers id.

const EutherpeTemplateNeedleURLSchema = "{{.URL-SCHEMA}}"
const EutherpeTemplateNeedleEutherpeAddr = "{{.EUTHERPE-ADDR}}"
const EutherpeTemplateNeedleEutherpe = "{{.EUTHERPE}}"
const EutherpeTemplateNeedleUpNext = "{{.UP-NEXT}}"
const EutherpeTemplateNeedleCollection = "{{.COLLECTION}}"
const EutherpeTemplateNeedlePlaylists = "{{.PLAYLISTS}}"
const EutherpeTemplateNeedleSelectedStorageDevice = "{{.SELECTED-STORAGE-DEVICE}}"
const EutherpeTemplateNeedleFoundStorageDevices = "{{.FOUND-STORAGE-DEVICES}}"
const EutherpeTemplateNeedleSelectedBluetoothDevice = "{{.SELECTED-BLUETOOTH-DEVICE}}"
const EutherpeTemplateNeedleFoundBluetoothDevices = "{{.FOUND-BLUETOOTH-DEVICES}}"
const EutherpeTemplateNeedleNowPlaying = "{{.NOW-PLAYING}}"
const EutherpeTemplateNeedleAlbumArtThumbnail = "{{.ALBUM-ART-THUMBNAIL}}"
const EutherpeTemplateNeedleLastError = "{{.LAST-ERROR}}"
const EutherpeTemplateNeedleRepeatAll = "{{.REPEAT-ALL}}"
const EutherpeTemplateNeedleRepeatOne = "{{.REPEAT-ONE}}"
const EutherpeTemplateNeedleCurrentConfig = "{{.CURRENT-CONFIG}}"
const EutherpeTemplateNeedleShuffleMode = "{{.SHUFFLE-MODE}}"
const EutherpeTemplateNeedlePlayMode = "{{.PLAY-MODE}}"
const EutherpeTemplateNeedlePlayerStatus = "{{.PLAYER-STATUS}}"
const EutherpeTemplateNeedleVolumeLevel = "{{.VOLUME-LEVEL}}"
const EutherpeTemplateNeedleCommonTags = "{{.COMMON-TAGS}}"
const EutherpeTemplateNeedleLastSelection = "{{.LAST-SELECTION}}"
const EutherpeTemplateNeedleAuthenticationModeSwitch = "{{.AUTHENTICATION-MODE-SWITCH}}"
const EutherpeTemplateNeedleHTTPSModeSwitch = "{{.HTTPS-MODE-SWITCH}}"
const EutherpeTemplateNeedleUpNextCount = "{{.UP-NEXT-COUNT}}"
const EutherpeTemplateNeedleFoundStorageDevicesCount = "{{.FOUND-STORAGE-DEVICES-COUNT}}"
const EutherpeTemplateNeedleFoundBluetoothDevicesCount = "{{.FOUND-BLUETOOTH-DEVICES-COUNT}}"

const EutherpeWebUIConfigSheetMusic = "Music"
const EutherpeWebUIConfigSheetCollection = "Collection"
const EutherpeWebUIConfigSheetPlaylists = "Playlists"
const EutherpeWebUIConfigSheetStorage = "Storage"
const EutherpeWebUIConfigSheetBluetooth = "Bluetooth"
const EutherpeWebUIConfigSheetSettings = "Settings"
const EutherpeWebUIConfigSheetDefault = EutherpeWebUIConfigSheetMusic

const EutherpeConfHome = "/etc/eutherpe"
const EutherpePlayerCache = "player.cache"
const EutherpePlaylistsHome = "playlists"
const EutherpeLastCollectionsHome = "collections"
const EutherpeCoversHome = "covers"
const EutherpeTagsHome = "tags"
