package vars

import (
    "internal/mplayer"
    "internal/dj"
    "internal/bluebraces"
    "sync"
    "os/exec"
)

type EutherpeVars struct {
    APPName string
    HTTPd struct {
        URLSchema string
        Addr string
    }
    BluetoothDevices []bluebraces.BluetoothDevice
    StorageDevices []string
    CachedDevices struct {
        BlueDevId string
        MusicDevId string
    }
    Collection mplayer.MusicCollection
    Playlists []dj.Playlist
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
    }
    Error error
    mtx sync.Mutex
}

func (e *EutherpeVars) Lock() {
    e.mtx.Lock()
}

func (e *EutherpeVars) Unlock() {
    e.mtx.Unlock()
}

func (e *EutherpeVars) Render(template string) string {
    return ""
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

// INFO(Rafael): Actions from "Collection" sheet.

const EutherpeCollectionAddSelectionToNextId = "collection-addselectiontonext"
const EutherpeCollectionAddSelectionToUpNextId = "collection-addselectiontoupnext"
const EutherpeCollectionAddSelectionToPlaylistId = "collection-addselectiontoplaylist"
const EutherpeCollectionTagSelectionAsId = "collection-tagselectionas"

// INFO(Rafael): Actions from "Playlists" sheet.

const EutherpePlaylistCreateId = "playlist-create"
const EutherpePlaylistRemoveId = "playlist-remove"
const EutherpePlaylistShowId = "playlist-show"
const EutherpePlaylistMoveUpId = "playlist-moveup"
const EutherpePlaylistMoveDownId = "playlist-movedown"
const EutherpePlaylistClearAllId = "playlist-clearall"
const EutherpePlaylistRemoveSongsId = "playlist-removesongs"

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

const EutherpePostFieldSelection = "selection"
const EutherpePostFieldPlaylist = "playlist"
const EutherpePostFieldStorageDevice = "storage-device"
const EutherpePostFieldBluetoothDevice = "bluetooth-device"

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
