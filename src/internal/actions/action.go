//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package actions

import (
    "internal/vars"
    "net/url"
    "net/http"
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
        case vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeMusicAddUpNextToPlaylistId:
            return AddSelectionToPlaylist
        case vars.EutherpeCollectionTagSelectionAsId:
            return TagSelection
        case vars.EutherpeCollectionUntagSelectionsId:
            return UntagSelection
        case vars.EutherpeCollectionPlayByGivenTagsId:
            return PlayByGivenTags
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
        case vars.EutherpePlaylistBackupPlaylistsId:
            return BackupPlaylists
        case vars.EutherpeStorageListId:
            return ListStorages
        case vars.EutherpeStorageScanId:
            return ScanStorage
        case vars.EutherpeStorageSetId:
            return SetStorage
        case vars.EutherpeStorageConvert2MP3Id:
            return Convert2MP3
        case vars.EutherpeStorageRestorePlaylistsId:
            return RestorePlaylists
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
        case vars.EutherpeSettingsFlickAuthModeId:
            return FlickAuthMode
        case vars.EutherpeSettingsChangePassphraseId:
            return ChangePassphrase
        case vars.EutherpeSettingsFlickHTTPSModeId:
            return FlickHTTPSMode
        case vars.EutherpeSettingsGenerateCertificateId:
            return GenerateCertificate
        case vars.EutherpeSettingsSetWLANCredentialsId:
            return SetWLANCredentials
        case vars.EutherpeSettingsSetHostNameId:
            return SetHostName
        case vars.EutherpeSettingsPowerOffId:
            return PowerOff
        case vars.EutherpeSettingsRebootId:
            return Reboot
        case vars.EutherpeSetCurrentConfigId:
            return SetCurrentConfig
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
             vars.EutherpeMusicAddUpNextToPlaylistId,
             vars.EutherpePlaylistReproduceId,
             vars.EutherpePlaylistReproduceSelectedOnesId,
             vars.EutherpeCollectionPlayByGivenTagsId,
             vars.EutherpeAuthenticateId:
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
             vars.EutherpePlaylistRemoveSongsId,
             vars.EutherpePlaylistBackupPlaylistsId:
            return vars.EutherpeWebUIConfigSheetPlaylists

        case vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeStorageConvert2MP3Id,
             vars.EutherpeStorageRestorePlaylistsId:
            return vars.EutherpeWebUIConfigSheetStorage

        case vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId:
            return vars.EutherpeWebUIConfigSheetBluetooth

        case vars.EutherpeSettingsFlickAuthModeId,
             vars.EutherpeSettingsChangePassphraseId,
             vars.EutherpeSettingsFlickHTTPSModeId,
             vars.EutherpeSettingsGenerateCertificateId,
             vars.EutherpeSettingsSetWLANCredentialsId,
             vars.EutherpeSettingsSetHostNameId,
             vars.EutherpeSettingsPowerOffId,
             vars.EutherpeSettingsRebootId:
            return vars.EutherpeWebUIConfigSheetSettings
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
             vars.EutherpeMusicAddUpNextToPlaylistId,
             vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpeCollectionUntagSelectionsId,
             vars.EutherpeCollectionPlayByGivenTagsId,
             vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpePlaylistRemoveSongsId,
             vars.EutherpePlaylistReproduceId,
             vars.EutherpePlaylistReproduceSelectedOnesId,
             vars.EutherpePlaylistBackupPlaylistsId,
             vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeStorageConvert2MP3Id,
             vars.EutherpeStorageRestorePlaylistsId,
             vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId,
             vars.EutherpeGetCommonTagsId,
             vars.EutherpeAuthenticateId,
             vars.EutherpeSettingsFlickAuthModeId,
             vars.EutherpeSettingsChangePassphraseId,
             vars.EutherpeSettingsFlickHTTPSModeId,
             vars.EutherpeSettingsGenerateCertificateId,
             vars.EutherpeSettingsSetWLANCredentialsId,
             vars.EutherpeSettingsSetHostNameId,
             vars.EutherpeSettingsPowerOffId,
             vars.EutherpeSettingsRebootId:
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
             vars.EutherpeMusicAddUpNextToPlaylistId,
             vars.EutherpeCollectionAddSelectionToNextId,
             vars.EutherpeCollectionAddSelectionToUpNextId,
             vars.EutherpeCollectionAddSelectionToPlaylistId,
             vars.EutherpeCollectionTagSelectionAsId,
             vars.EutherpeCollectionUntagSelectionsId,
             vars.EutherpeCollectionPlayByGivenTagsId,
             vars.EutherpePlaylistCreateId,
             vars.EutherpePlaylistRemoveId,
             vars.EutherpePlaylistShowId,
             vars.EutherpePlaylistMoveUpId,
             vars.EutherpePlaylistMoveDownId,
             vars.EutherpePlaylistClearAllId,
             vars.EutherpePlaylistRemoveSongsId,
             vars.EutherpePlaylistReproduceId,
             vars.EutherpePlaylistReproduceSelectedOnesId,
             vars.EutherpePlaylistBackupPlaylistsId,
             vars.EutherpeStorageListId,
             vars.EutherpeStorageScanId,
             vars.EutherpeStorageSetId,
             vars.EutherpeStorageConvert2MP3Id,
             vars.EutherpeStorageRestorePlaylistsId,
             vars.EutherpeBluetoothProbeDevicesId,
             vars.EutherpeBluetoothPairId,
             vars.EutherpeBluetoothUnpairId,
             vars.EutherpeBluetoothTrustId,
             vars.EutherpeBluetoothUntrustId,
             vars.EutherpeGetCommonTagsId,
             vars.EutherpeSettingsFlickAuthModeId,
             vars.EutherpeSettingsChangePassphraseId,
             vars.EutherpeSettingsFlickHTTPSModeId,
             vars.EutherpeSettingsGenerateCertificateId,
             vars.EutherpeSettingsSetWLANCredentialsId,
             vars.EutherpeSettingsSetHostNameId,
             vars.EutherpeSettingsPowerOffId,
             vars.EutherpeSettingsRebootId,
             vars.EutherpeSetCurrentConfigId:
            eutherpeVars.HTTPd.ResponseWriter.WriteHeader(http.StatusOK)
            return eutherpeVars.HTTPd.IndexHTML

        case vars.EutherpeAuthenticateId:
            eutherpeVars.HTTPd.ResponseWriter.WriteHeader(http.StatusOK)
            if eutherpeVars.LastError == nil {
                return "<html><script>window.location=\"{{.URL-SCHEMA}}://{{.EUTHERPE-ADDR}}/eutherpe\"</script></html>"
            } else {
                return eutherpeVars.HTTPd.LoginHTML
            }
            break

        case vars.EutherpePlayerStatusId:
            eutherpeVars.HTTPd.ResponseWriter.WriteHeader(http.StatusOK)
            return vars.EutherpeTemplateNeedlePlayerStatus
    }
    eutherpeVars.HTTPd.ResponseWriter.WriteHeader(http.StatusNotFound)
    return eutherpeVars.HTTPd.ErrorHTML
}
