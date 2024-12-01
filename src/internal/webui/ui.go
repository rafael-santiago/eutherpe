//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package webui

import (
    "internal/vars"
    "internal/actions"
    "internal/renders"
    "internal/bluebraces"
    "fmt"
    "os"
    "path"
    //"path/filepath"
    "os/signal"
    "syscall"
    "net/http"
    "net/url"
    "sync"
    "strings"
    "time"
    "context"
)

type EutherpeHTTPHandler struct {
    requestTurnstile sync.Mutex
    eutherpeVars *vars.EutherpeVars
}

func RunWebUI(eutherpeVars *vars.EutherpeVars) error {
    if eutherpeVars == nil {
        fmt.Errorf("panic: nil eutherpeVars!\n")
    }
    if eutherpeVars.Player.Shuffle {
        eutherpeVars.Player.Shuffle = false
        actions.MusicShuffle(eutherpeVars, nil)
    }
    sigintWatchdog := make(chan os.Signal, 1)
    signal.Notify(sigintWatchdog, os.Interrupt)
    signal.Notify(sigintWatchdog, syscall.SIGINT, syscall.SIGTERM)

    var err error
    eutherpeMiniplayer := &http.Server{}
    go eutherpeHTTPd(EutherpeHTTPHandler { sync.Mutex{}, eutherpeVars },
                     sigintWatchdog, eutherpeMiniplayer, &err)
    goinHome := make(chan bool, 1)
    go func() {
        const kSavingWindow = 42
        timeToSave := time.Now().Unix() + kSavingWindow
        for {
            select {
                case <-goinHome:
                    break
                default:
                    if (time.Now().Unix() >= timeToSave) {
                        eutherpeVars.Lock()
                        eutherpeVars.SaveSession()
                        eutherpeVars.Unlock()
                        timeToSave = time.Now().Unix() + kSavingWindow
                    }
                    time.Sleep(1 * time.Second)
            }
        }
    }()
    <-sigintWatchdog
    goinHome <- true
    if eutherpeMiniplayer != nil {
        shutdownStatus := eutherpeMiniplayer.Shutdown(context.Background())
        if shutdownStatus != nil {
            fmt.Fprintf(os.Stderr,
                        "error: when trying to shutdown web miniplayer : '%s'\n", shutdownStatus.Error())
        } else {
            fmt.Fprintf(os.Stdout, "info: web miniplayer has exit.\n")
            err = nil
        }
    }
    // INFO(Rafael): It is important otherwise Eutherpe can exits by letting music playing
    //               sometimes.
    eutherpeVars.Player.AutoPlay = !eutherpeVars.Player.Stopped
    actions.MusicStop(eutherpeVars, nil)
    if len(eutherpeVars.CachedDevices.BlueDevId) > 0 {
        bluebraces.DisconnectDevice(eutherpeVars.CachedDevices.BlueDevId)
        bluebraces.UnpairDevice(eutherpeVars.CachedDevices.BlueDevId)
    }
    return err
}

func eutherpeHTTPd(eutherpeHTTPHandler EutherpeHTTPHandler,
                   sigintWatchdog chan os.Signal,
                   eutherpeMiniplayer *http.Server,
                   err *error) {
    http.HandleFunc("/", eutherpeHTTPHandler.handler)
    eutherpeMiniplayer.Addr = eutherpeHTTPHandler.eutherpeVars.HTTPd.Addr + ":" +
                              eutherpeHTTPHandler.eutherpeVars.HTTPd.Port
    if !eutherpeHTTPHandler.eutherpeVars.HTTPd.TLS {
        (*err) = eutherpeMiniplayer.ListenAndServe()
    } else {
        cerFilePath := path.Join(eutherpeHTTPHandler.eutherpeVars.HTTPd.PubRoot, "cert/eutherpe.cer")
        privKeyFilePath := path.Join(eutherpeHTTPHandler.eutherpeVars.ConfHome, "eutherpe.priv")
        (*err) = eutherpeMiniplayer.ListenAndServeTLS(cerFilePath, privKeyFilePath)
    }
    if (*err) != nil {
        fmt.Fprintf(os.Stderr, "panic: %s\n", (*err).Error())
        sigintWatchdog <- syscall.SIGINT
        eutherpeMiniplayer = nil
    }
}

func (ehh *EutherpeHTTPHandler) handler(w http.ResponseWriter, r *http.Request) {
    ehh.requestTurnstile.Lock()
    defer ehh.requestTurnstile.Unlock()
    var templatedOutput string
    if len(ehh.eutherpeVars.HostName) > 0 &&
       strings.HasPrefix(r.Host, ehh.eutherpeVars.HostName) {
        ehh.eutherpeVars.HTTPd.RequestedByHostName = true
    } else {
        ehh.eutherpeVars.HTTPd.RequestedByHostName = false
    }
    var template uint = vars.EutherpeNoTemplate
    switch r.URL.Path {
        case "/eutherpe":
                var contentType = "text/html"
                if r.Method == "GET" {
                    if ehh.eutherpeVars.HTTPd.Authenticated &&
                      !ehh.eutherpeVars.HTTPd.AuthWatchdog.IsAuthenticated(r.RemoteAddr) {
                        templatedOutput = ehh.eutherpeVars.HTTPd.LoginHTML
                        template = vars.EutherpeGateTemplate
                    } else {
                        templatedOutput = ehh.eutherpeVars.HTTPd.IndexHTML
                        template = vars.EutherpeIndexTemplate
                        // INFO(Rafael): It is needed otherwise we will get a redirection loop.
                        ehh.eutherpeVars.RenderedIndexHTML = ""
                    }
                    if len(ehh.eutherpeVars.CurrentConfig) == 0 {
                        ehh.eutherpeVars.CurrentConfig = vars.EutherpeWebUIConfigSheetDefault
                    }
                } else if r.Method == "POST" {
                    r.ParseForm()
                    if ehh.eutherpeVars.HTTPd.Authenticated &&
                       !ehh.eutherpeVars.HTTPd.AuthWatchdog.IsAuthenticated(r.RemoteAddr) {
                        templatedOutput = ehh.eutherpeVars.HTTPd.LoginHTML
                        template = vars.EutherpeGateTemplate
                    } else {
                        ehh.eutherpeVars.HTTPd.AuthWatchdog.RefreshAuthWindow(r.RemoteAddr)
                        actionHandler := actions.GetEutherpeActionHandler(&r.Form)
                        templatedOutput = ehh.processAction(actionHandler, &r.Form)
                        // INFO(Rafael): It is needed to force a re-caching.
                        ehh.eutherpeVars.RenderedIndexHTML = ""
                    }
                    contentType = actions.GetContentTypeByActionId(&r.Form)
                } else {
                    templatedOutput = ehh.eutherpeVars.HTTPd.ErrorHTML
                    ehh.eutherpeVars.LastError = fmt.Errorf("501 Not Implemented")
                }
                w.Header().Set("content-type", contentType)
            break
        case "/eutherpe-auth":
            if !ehh.eutherpeVars.HTTPd.Authenticated {
                templatedOutput = ehh.eutherpeVars.HTTPd.ErrorHTML
                ehh.eutherpeVars.LastError = fmt.Errorf("303 See Other")
            } else if r.Method == "GET" {
                templatedOutput = ehh.eutherpeVars.HTTPd.LoginHTML
                template = vars.EutherpeIndexTemplate
            } else if r.Method == "POST" {
                r.ParseForm()
                r.Form.Add(vars.EutherpePostFieldRemoteAddr, r.RemoteAddr)
                templatedOutput = ehh.processAction(actions.Authenticate, &r.Form)
                r.URL.Path = "/eutherpe"
                template = vars.EutherpeIndexTemplate
            } else {
                templatedOutput = ehh.eutherpeVars.HTTPd.ErrorHTML
                ehh.eutherpeVars.LastError = fmt.Errorf("501 Not Implemented")
            }
            w.Header().Set("content-type", "text/html")
            break
        default:
            templatedOutput = ehh.processGET(&w, r)
    }
    fmt.Fprintf(w, "%s", renders.RenderData(templatedOutput, ehh.eutherpeVars, template))
    ehh.eutherpeVars.LastError = nil
}

func (ehh *EutherpeHTTPHandler) processAction(actionHandler actions.EutherpeActionFunc, userData *url.Values) string {
    if actionHandler == nil {
        ehh.eutherpeVars.LastError =
                fmt.Errorf("500 Internal Server Error")
        return actions.GetVDocByActionId(userData, ehh.eutherpeVars)
    }
    ehh.eutherpeVars.CurrentConfig = actions.CurrentConfigByActionId(userData)
    ehh.eutherpeVars.LastError = actionHandler(ehh.eutherpeVars, userData)
    if ehh.eutherpeVars.LastError != nil {
        userData.Del(vars.EutherpePostFieldLastError)
        userData.Add(vars.EutherpePostFieldLastError, ehh.eutherpeVars.LastError.Error())
    }
    return actions.GetVDocByActionId(userData, ehh.eutherpeVars)
}

func (ehh *EutherpeHTTPHandler) processGET(w *http.ResponseWriter, r *http.Request) string {
    vdoc := r.URL.Path
    if !ehh.isPubFile(vdoc) {
        ehh.eutherpeVars.LastError = fmt.Errorf("403 Forbidden")
        return ehh.eutherpeVars.HTTPd.ErrorHTML
    }
    (*w).Header().Set("content-type", GetMIMEType(vdoc))
    data, err := os.ReadFile(path.Join(ehh.eutherpeVars.HTTPd.PubRoot, vdoc))
    if err != nil {
        return "(null)"
    }
    return string(data)
}

func (ehh *EutherpeHTTPHandler) isPubFile(wantedFilePath string) bool {
    for _, pubFile := range ehh.eutherpeVars.HTTPd.PubFiles {
        if wantedFilePath == pubFile {
            return true
        }
    }
    return false
}
