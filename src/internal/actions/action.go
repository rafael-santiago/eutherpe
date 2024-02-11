package actions

import (
    "internal/vars"
    "net/url"
)

type EutherpeActionFunc func(eutherpeInstance *vars.EutherpeVars, userData *url.Values) error

func GetEutherpeActionHandler(userData *url.Values) EutherpeActionFunc {
    switch userData.Get(vars.EutherpeActionId) {
        case vars.EutherpeMusicRemoveId:
            return MusicRemove
        case vars.EutherpeMusicMoveUpId:
            return MusicMoveUp
        case vars.EutherpeMusicMoveDownId:
            return MusicMoveDown
        case vars.EutherpeMusicClearAllId:
            return MusicClearAll
        case vars.EutherpeMusicShuffleId:
            return MusicShuffle
        case vars.EutherpeMusicRepeatAllId:
            return MusicRepeatAll
        case vars.EutherpeMusicRepeatOneId:
            return MusicRepeatOne
        case vars.EutherpeMusicPlayId:
            return MusicPlay
        case vars.EutherpeMusicStopId:
            return MusicStop
        case vars.EutherpeMusicNextId:
            return MusicNext
        case vars.EutherpeCollectionAddSelectionToNextId:
            return AddSelectionToNext
        case vars.EutherpeCollectionAddSelectionToUpNextId:
            return AddSelectionToUpNext
        case vars.EutherpeCollectionAddSelectionToPlaylistId:
            return AddSelectionToPlaylist
        case vars.EutherpeCollectionTagSelectionAsId:
            return nil
        case vars.EutherpePlaylistCreateId:
            return CreatePlaylist
        case vars.EutherpePlaylistRemoveId:
            return RemovePlaylist
        case vars.EutherpePlaylistShowId:
            return ShowPlaylist
        case vars.EutherpePlaylistMoveUpId:
            return MoveUpPlaylistSongs
        case vars.EutherpePlaylistMoveDownId:
            return MoveDownPlaylistSongs
        case vars.EutherpePlaylistClearAllId:
            return ClearAllPlaylist
        case vars.EutherpeStorageListId:
            return ListStorages
        case vars.EutherpeStorageScanId:
            return ScanStorage
        case vars.EutherpeStorageSetId:
            return SetStorage
        case vars.EutherpeBluetoothProbeDevicesId:
            return ProbeBluetoothDevices
        case vars.EutherpeBluetoothPairId:
            return PairBluetoothDevice
        case vars.EutherpeBluetoothUnpairId:
            return UnpairBluetoothDevice
        case vars.EutherpeBluetoothTrustId:
            return TrustBluetoothDevice
        case vars.EutherpeBluetoothUntrustId:
            return UntrustBluetoothDevice
    }
    return nil
}

func CurrentConfigByActionId(userData *url.Values) string {
    switch userData.Get(vars.EutherpeActionId) {
        case vars.EutherpeMusicRemoveId,
             vars.EutherpeMusicMoveUpId,
             vars.EutherpeMusicMoveDownId,
             vars.EutherpeMusicClearAllId,
             vars.EutherpeMusicShuffleId,
             vars.EutherpeMusicRepeatAllId,
             vars.EutherpeMusicRepeatOneId,
             vars.EutherpeMusicPlayId,
             vars.EutherpeMusicNextId,
             vars.EutherpeMusicStopId:
            return vars.EutherpeWebUIConfigSheetMusic

        case vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId:
            return vars.EutherpeWebUIConfigSheetCollection

        case vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId:
            return vars.EutherpeWebUIConfigSheetPlaylists

        case vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId:
            return vars.EutherpeWebUIConfigSheetStorage

        case vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId:
            return vars.EutherpeWebUIConfigSheetBluetooth
    }
    return vars.EutherpeWebUIConfigSheetDefault
}

func GetContentTypeByActionId(userData *url.Values) string {
    switch userData.Get(vars.EutherpeActionId) {
        case vars.EutherpeMusicRemoveId,
             vars.EutherpeMusicMoveUpId,
             vars.EutherpeMusicMoveDownId,
             vars.EutherpeMusicClearAllId,
             vars.EutherpeMusicShuffleId,
             vars.EutherpeMusicRepeatAllId,
             vars.EutherpeMusicRepeatOneId,
             vars.EutherpeMusicPlayId,
             vars.EutherpeMusicNextId,
             vars.EutherpeMusicStopId,
             vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId:
            return "text/html"
    }
    return "application/json"
}

func GetVDocByActionId(userData *url.Values, eutherpeVars *vars.EutherpeVars) string {
    switch userData.Get(vars.EutherpeActionId) {
        case vars.EutherpeMusicRemoveId,
             vars.EutherpeMusicMoveUpId,
             vars.EutherpeMusicMoveDownId,
             vars.EutherpeMusicClearAllId,
             vars.EutherpeMusicShuffleId,
             vars.EutherpeMusicRepeatAllId,
             vars.EutherpeMusicRepeatOneId,
             vars.EutherpeMusicPlayId,
             vars.EutherpeMusicNextId,
             vars.EutherpeMusicStopId,
             vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId:
            return eutherpeVars.HTTPd.IndexHTML

        case vars.EutherpePlayerStatusId:
            return vars.EutherpeTemplateNeedlePlayerStatus
    }

    return eutherpeVars.HTTPd.ErrorHTML
}
