package actions

import (
    "internal/vars"
    "net/url"
)

type EutherpeActionFunc func(eutherpeInstance *vars.EutherpeVars, userData *url.Values) error

func GetEutherpeActionHandler(userData *url.Values) EutherpeActionFunc {
    switch userData.Get(vars.EutherpeActionId) {
        case vars.EutherpeMusicRemoveId:
            return nil
        case vars.EutherpeMusicMoveUpId:
            return nil
        case vars.EutherpeMusicMoveDownId:
            return nil
        case vars.EutherpeMusicClearAllId:
            return nil
        case vars.EutherpeMusicShuffleId:
            return nil
        case vars.EutherpeMusicRepeatAllId:
            return nil
        case vars.EutherpeMusicRepeatOneId:
            return nil
        case vars.EutherpeMusicPlayId:
            return nil
        case vars.EutherpeMusicNextId:
            return nil
        case vars.EutherpeCollectionAddSelectionToNextId:
            return nil
        case vars.EutherpeCollectionAddSelectionToUpNextId:
            return nil
        case vars.EutherpeCollectionAddSelectionToPlaylistId:
            return nil
        case vars.EutherpeCollectionTagSelectionAsId:
            return nil
        case vars.EutherpePlaylistCreateId:
            return nil
        case vars.EutherpePlaylistRemoveId:
            return nil
        case vars.EutherpePlaylistShowId:
            return nil
        case vars.EutherpePlaylistMoveUpId:
            return nil
        case vars.EutherpePlaylistMoveDownId:
            return nil
        case vars.EutherpePlaylistClearAllId:
            return nil
        case vars.EutherpeStorageListId:
            return nil
        case vars.EutherpeStorageScanId:
            return nil
        case vars.EutherpeStorageSetId:
            return nil
        case vars.EutherpeBluetoothProbeDevicesId:
            return nil
        case vars.EutherpeBluetoothPairId:
            return nil
        case vars.EutherpeBluetoothUnpairId:
            return nil
        case vars.EutherpeBluetoothTrustId:
            return nil
        case vars.EutherpeBluetoothUntrustId:
            return nil
    }
    return nil
}
