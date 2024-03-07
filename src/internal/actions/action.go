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
        case vars.EutherpeMusicLastId:
            return MusicLast
        case vars.EutherpeMusicSetVolumeId:
            return MusicSetVolume
        case vars.EutherpeCollectionAddSelectionToNextId:
            return AddSelectionToNext
        case vars.EutherpeCollectionAddSelectionToUpNextId:
            return AddSelectionToUpNext
        case vars.EutherpeCollectionAddSelectionToPlaylistId:
            return AddSelectionToPlaylist
        case vars.EutherpeCollectionTagSelectionAsId:
            return TagSelection
        case vars.EutherpeCollectionUntagSelectionsId:
            return UntagSelection
        case vars.EutherpeGetCommonTagsId:
            return GetCommonTags
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
        case vars.EutherpePlaylistRemoveSongsId:
            return RemoveSongsFromPlaylist
        case vars.EutherpePlaylistReproduceId:
            return ReproducePlaylist
        case vars.EutherpePlaylistReproduceSelectedOnesId:
            return ReproduceSelectedOnesFromPlaylist
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
             vars.EutherpeMusicLastId,
             vars.EutherpeMusicStopId,
             vars.EutherpeMusicSetVolumeId,
             vars.EutherpePlaylistReproduceId,
             vars.EutherpePlaylistReproduceSelectedOnesId:
            return vars.EutherpeWebUIConfigSheetMusic

        case vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpeCollectionUntagSelectionsId,
             vars.EutherpeGetCommonTagsId:
            return vars.EutherpeWebUIConfigSheetCollection

        case vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpePlaylistRemoveSongsId:
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
             vars.EutherpeMusicLastId,
             vars.EutherpeMusicStopId,
             vars.EutherpeMusicSetVolumeId,
             vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpeCollectionUntagSelectionsId,
             vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpePlaylistRemoveSongsId,
             vars.EutherpePlaylistReproduceId,
             vars.EutherpePlaylistReproduceSelectedOnesId,
             vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId,
             vars.EutherpeGetCommonTagsId:
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
             vars.EutherpeMusicLastId,
             vars.EutherpeMusicStopId,
             vars.EutherpeMusicSetVolumeId,
             vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpeCollectionUntagSelectionsId,
             vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpePlaylistRemoveSongsId,
             vars.EutherpePlaylistReproduceId,
             vars.EutherpePlaylistReproduceSelectedOnesId,
             vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId,
             vars.EutherpeGetCommonTagsId:
            return eutherpeVars.HTTPd.IndexHTML

        //case vars.EutherpeGetCommonTagsId:
        //    return vars.EutherpeTemplateNeedleCommonTags

        case vars.EutherpePlayerStatusId:
            return vars.EutherpeTemplateNeedlePlayerStatus
    }

    return eutherpeVars.HTTPd.ErrorHTML
}
