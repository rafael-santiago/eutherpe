package renders

import (
    "internal/vars"
)

type EutherpeDataRendererFunc func(templatedInput string, eutherpeVars *vars.EutherpeVars) string

func RenderData(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var doRenderFuncs []EutherpeDataRendererFunc = []EutherpeDataRendererFunc {
        AlbumArtThumbnailRender, CollectionRender, EutherpeAddrRender,
        EutherpeRender, FoundBluetoothDevicesRender, FoundStorageDevicesRender,
        NowPlayingRender, PlaylistsRender, SelectedBluetoothDeviceRender,
        SelectedStorageDeviceRender, UpNextRender, URLSchemaRender,
        LastErrorRender, RepeatAllRender, RepeatOneRender, CurrentConfigRender,
        ShuffleModeRender, PlayModeRender, PlayerStatusRender, VolumeLevelRender,
        CommonTagsRender, LastSelectionRender, AuthenticationModeSwitchRender,
        UpNextCountRender, FoundStorageDevicesCountRender, FoundBluetoothDevicesCountRender,
        HTTPSModeSwitchRender, ESSIDRender, HostNameRender,
    }
    var output string = templatedInput
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    for _, doRender := range doRenderFuncs {
        output = doRender(output, eutherpeVars)
    }
    return output
}
