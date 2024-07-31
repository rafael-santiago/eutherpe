//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
)

type EutherpeDataRendererFunc func(templatedInput string, eutherpeVars *vars.EutherpeVars) string

func RenderData(templatedInput string, eutherpeVars *vars.EutherpeVars, template ...uint) string {
    var doRenderFuncs []EutherpeDataRendererFunc = []EutherpeDataRendererFunc {
        AlbumArtThumbnailRender, CollectionRender, EutherpeAddrRender,
        EutherpeRender, FoundBluetoothDevicesRender, FoundStorageDevicesRender,
        NowPlayingRender, PlaylistsRender, SelectedBluetoothDeviceRender,
        SelectedStorageDeviceRender, UpNextRender, URLSchemaRender,
        LastErrorRender, RepeatAllRender, RepeatOneRender, CurrentConfigRender,
        ShuffleModeRender, PlayModeRender, PlayerStatusRender, VolumeLevelRender,
        CommonTagsRender, LastSelectionRender, AuthenticationModeSwitchRender,
        UpNextCountRender, FoundStorageDevicesCountRender, FoundBluetoothDevicesCountRender,
        HTTPSModeSwitchRender, ESSIDRender, HostNameRender, VersionRender, CopyrightRender,
    }
    if len(template) > 0 {
        if template[0] == vars.EutherpeIndexTemplate &&
            len(eutherpeVars.RenderedIndexHTML) > 0 {
            return eutherpeVars.RenderedIndexHTML
        } else if template[0] == vars.EutherpeGateTemplate &&
                   len(eutherpeVars.RenderedGateHTML) > 0 {
            return eutherpeVars.RenderedGateHTML
        }
    }
    var output string = templatedInput
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    for _, doRender := range doRenderFuncs {
        output = doRender(output, eutherpeVars)
    }
    if len(template) > 0 {
        if template[0] == vars.EutherpeIndexTemplate {
            eutherpeVars.RenderedIndexHTML = output
        } else if template[0] == vars.EutherpeGateTemplate {
            eutherpeVars.RenderedGateHTML = output
        }
    }
    return output
}
