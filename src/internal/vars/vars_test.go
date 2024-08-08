//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package vars

import (
    "testing"
)

func Test_shoulSaveSession(t *testing.T) {
    eutherpeVars := &EutherpeVars{}
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.HTTPd.Authenticated = true
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.HTTPd.TLS = true
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.CachedDevices.BlueDevId = "blau one two three"
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.CachedDevices.MusicDevId = "segura o tcham!"
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.CollectionHTML = "ziriguidum"
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.UpNextHTML = "seria melhor ter ido ver o filme do pele..."
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.PlaylistsHTML = "embaixo desse morro passa boi, passa boiada, olha so que movimento"
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.RenderedIndexHTML = "domingo eu venho passar a segunda-feira com voces..."
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.RenderedGateHTML = "chega uma hora que eu fico sem ideia para bobeira"
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.RenderedAlbumArtThumbnailHTML = "mas nem assim tomo jeito e, logo a bobeira volta..."
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.Player.Shuffle = true
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.Player.RepeatAll = true
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.Player.RepeatOne = true
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.Player.VolumeLevel = 42
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.WLAN.ESSID = "pois essas ideias, como as concebo (hmmmm que dilicia)... esquece! so queria dizer que uma vez conheci uma mulher de goias que confisca gado... ela era poliglota e me disse que catalogado inclusive e o jeito que o vaqueiro espanhol diz que esta indo pegar um boi... yo voy catalogado!... pronto, chega!... sou um oceano de cultura (inutil)."
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
    eutherpeVars.Player.UpNextCurrentOffset = 10
    if !eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is not detecting changes accordingly.")
    }
    if eutherpeVars.shouldSaveSession() {
        t.Errorf("shouldSaveSession() is detecting changes that did not happen.")
    }
}
