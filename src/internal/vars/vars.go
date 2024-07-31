//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package vars

import (
    "internal/mplayer"
    "internal/dj"
    "internal/bluebraces"
    "internal/storage"
    "internal/auth"
    "internal/wifi"
    "internal/mdns"
    "sync"
    "os/exec"
    "runtime"
    "encoding/json"
    "encoding/base64"
    "os"
    "path"
    "fmt"
    "strings"
    "net"
    "time"
)

type EutherpeVars struct {
    APPName string
    HostName string
    ConfHome string
    HTTPd struct {
        Authenticated bool
        TLS bool
        AuthWatchdog *auth.AuthWatchdog
        HashKey string
        URLSchema string
        Addr string
        Port string
        PubRoot string
        PubFiles []string
        IndexHTML string
        ErrorHTML string
        LoginHTML string
        RequestedByHostName bool
    }
    MDNS struct {
        Hosts []mdns.MDNSHost
        GoinHome chan bool
    }
    BluetoothDevices []bluebraces.BluetoothDevice
    StorageDevices []string
    CachedDevices struct {
        BlueDevId string
        MusicDevId string
        MixerControlName string
    }
    Collection mplayer.MusicCollection
    CollectionHTML string
    UpNextHTML string
    PlaylistsHTML string
    RenderedIndexHTML string
    RenderedGateHTML string
    RenderedAlbumArtThumbnailHTML string
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
    WLAN struct {
        ESSID string
        Iface string
        ConnSession *exec.Cmd
        Addr string
    }
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
    ESSID string
    HostName string
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
                                         e.HTTPd.TLS,
                                         e.WLAN.ESSID,
                                         e.HostName, }
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
    e.WLAN.ESSID = cachedData.ESSID
    e.HostName = cachedData.HostName
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

func (e *EutherpeVars) SetAddr() error {
    ifaces, err := net.Interfaces()
    if err != nil {
        return err
    }
    for _, iface := range ifaces {
        if (iface.Flags & net.FlagLoopback) == 0 {
            addrs, err := iface.Addrs()
            if err != nil {
                continue
            }
            for _, addr := range addrs {
                ip, _, err := net.ParseCIDR(addr.String())
                if err == nil {
                    e.HTTPd.Addr = ip.String()
                    break
                }
            }
            if len(e.HTTPd.Addr) > 0 {
                break
            }
        }
    }
    if len(e.WLAN.ESSID) > 0 {
        fmt.Printf("info: WLAN is configured, trying to acquire a WLAN connection... wait...\n")
        wlanIfaces := wifi.GetIfaces()
        if len(wlanIfaces) == 0 {
            fmt.Printf("warn: No WLAN interface was found.\n")
        } else {
            wifi.SetIfaceUp(wlanIfaces[0])
            e.WLAN.ConnSession, err = wifi.Start(wlanIfaces[0])
            if err == nil {
                e.WLAN.Iface = wlanIfaces[0]
                ipAddr, _ := wifi.LeaseAddr(wlanIfaces[0])
                if len(ipAddr) == 0 {
                    wifi.Stop(e.WLAN.ConnSession)
                    e.WLAN.ConnSession = nil
                } else {
                    e.HTTPd.Addr = ipAddr
                    fmt.Printf("info: Eutherpe has ingressed to the WLAN %s.\n", e.WLAN.ESSID)
                }
            }
        }
    }
    if len(e.HTTPd.Addr) == 0 {
        return fmt.Errorf("Unable to set a valid IP")
    }
    return nil
}

func (e *EutherpeVars) getConfHome() string {
    st, err := os.Stat("/etc/eutherpe")
    if err == nil && st.IsDir() {
        return "/etc/eutherpe"
    }
    cwd, err := os.Getwd()
    if err != nil {
        return ""
    }
    localEtcEutherpe := path.Join(cwd, "etc", "eutherpe")
    st, err = os.Stat(localEtcEutherpe)
    if err == nil && st.IsDir() {
        return localEtcEutherpe
    }
    return ""
}

func (e *EutherpeVars) getPubRoot() string {
    pubRootDirPath := path.Join(e.ConfHome, "web")
    st, err := os.Stat(pubRootDirPath)
    if err == nil && st.IsDir() {
        return pubRootDirPath
    }
    pubRootDirPath, err = os.Getwd()
    if err != nil {
        return ""
    }
    pubRootDirPath = path.Join(pubRootDirPath, "web")
    st, err = os.Stat(pubRootDirPath)
    if err == nil && st.IsDir() {
        return pubRootDirPath
    }
    return ""
}

func (e *EutherpeVars) setEutherpePubTrinket() []string {
    pubFiles := make([]string, 0)
    pubFiles = append(pubFiles, "/js/eutherpe.js")
    pubFiles = append(pubFiles, "/css/eutherpe.css")
    pubFiles = append(pubFiles, "/fonts/Sabo-Filled.otf")
    pubFiles = append(pubFiles, "/fonts/Sabo-Regular.otf")
    pubFiles = append(pubFiles, "/cert/eutherpe.cer")
    return pubFiles
}

func (e *EutherpeVars) TuneUp() {
    e.ConfHome = e.getConfHome()
    if len(e.ConfHome) == 0 {
        fmt.Fprintf(os.Stderr, "error: unable to found out Eutherpe's config folder.\n")
        os.Exit(1)
    }
    e.Player.RepeatAll = false
    e.Player.RepeatOne = false
    e.Player.Stopped = true
    if e.Player.VolumeLevel == 0 {
        e.Player.VolumeLevel = mplayer.GetVolumeLevel(false)
    }
    e.HTTPd.URLSchema = "http"
    e.HTTPd.PubRoot = e.getPubRoot()
    e.HTTPd.PubFiles = e.setEutherpePubTrinket()
    data, err := os.ReadFile(path.Join(e.HTTPd.PubRoot, "html", "eutherpe.html"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "i/o error: '%s'\n", err.Error())
        os.Exit(1)
    }
    e.HTTPd.IndexHTML = string(data)
    data, err = os.ReadFile(path.Join(e.HTTPd.PubRoot, "html", "error.html"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "i/o error: '%s'\n", err.Error())
        os.Exit(1)
    }
    e.HTTPd.ErrorHTML = string(data)
    data, err = os.ReadFile(path.Join(e.HTTPd.PubRoot, "html", "eutherpe-gate.html"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "i/o error: '%s'\n", err.Error())
        os.Exit(1)
    }
    e.HTTPd.LoginHTML = string(data)
    e.HTTPd.AuthWatchdog = auth.NewAuthWatchdog(time.Duration(15 * time.Minute))
    e.HTTPd.AuthWatchdog.On()
    e.RestoreSession()
    e.SetAddr()
    if e.WLAN.ConnSession != nil {
        defer wifi.ReleaseAddr(e.WLAN.Iface)
        defer wifi.Stop(e.WLAN.ConnSession)
        defer wifi.SetIfaceDown(e.WLAN.Iface)
    }
    if strings.HasPrefix(runtime.GOARCH, "arm") && len(e.HostName) == 0 {
        // INFO(Rafael): It is convenient because find out ip address of a
        //               raspberry pi it is a pain in the neck.
        e.HostName = "eutherpe.local"
    }
    if len(e.HostName) > 0 {
        e.MDNS.GoinHome = make(chan bool)
        e.MDNS.Hosts = make([]mdns.MDNSHost, 0)
        ipAddr := net.ParseIP(e.HTTPd.Addr)
        if strings.Index(e.HTTPd.Addr, ".") > - 1 {
            ipAddr = ipAddr[12:16]
        }
        e.MDNS.Hosts = append(e.MDNS.Hosts, mdns.MDNSHost { e.HostName, ipAddr, 300, })
        go mdns.MDNSServerStart(e.MDNS.Hosts, e.MDNS.GoinHome)
    }
    e.HTTPd.Port = "8080"
}

const EutherpeVersion = "v1"
const EutherpeCopyrightDisclaimer = "Eutherpe is Copyright (c) 2024 by Rafael Santiago<br><br>You can redistribute it and/or modify under the terms of the GNU General Public License version 2.<br><br>Bug reports, feedback etc: <a href=\"mailto:voidbrainvoid@tutanota.com\"?subject=\"[Eutherpe] <here goes the subject>\"><u>mail me</u></a> or open an <a href=\"https://github.com/rafael-santiago/eutherpe/issues\"><u>issue</u></a> at Eutherpe's project repository. Thanks in advance! &#x1F609"

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
const EutherpeStorageConvert2MP3Id = "convert-2mp3"

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
const EutherpeSettingsSetWLANCredentialsId = "settings-setwlancredentials"
const EutherpeSettingsSetHostNameId = "settings-sethostname"
const EutherpeSettingsPowerOffId = "settings-poweroff"
const EutherpeSettingsRebootId = "settings-reboot"

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
const EutherpePostFieldESSID = "essid"
const EutherpePostFieldHostName = "hostname"

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
const EutherpeTemplateNeedleESSID = "{{.ESSID}}"
const EutherpeTemplateNeedleHostName = "{{.HOSTNAME}}"
const EutherpeTemplateNeedleVersion = "{{.EUTHERPE-VERSION}}"
const EutherpeTemplateNeedleCopyrightDisclaimer = "{{.EUTHERPE-COPYRIGHT-DISCLAIMER}}"

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

const EutherpeNoTemplate = 0
const EutherpeIndexTemplate = 1
const EutherpeGateTemplate = 2
