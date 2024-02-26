package webui

import (
    "internal/vars"
    "internal/actions"
    "internal/renders"
    "fmt"
    "os"
    "path"
    //"path/filepath"
    "os/signal"
    "syscall"
    "net/http"
    "net/url"
)

type EutherpeHTTPHandler struct {
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
    var err error = nil
    sigintWatchdog := make(chan os.Signal, 1)
    go eutherpeHTTPd(EutherpeHTTPHandler { eutherpeVars }, sigintWatchdog, &err)
    signal.Notify(sigintWatchdog, os.Interrupt)
    signal.Notify(sigintWatchdog, syscall.SIGINT|syscall.SIGTERM)
    <-sigintWatchdog
    return err
}

func eutherpeHTTPd(eutherpeHTTPHandler EutherpeHTTPHandler, sigintWatchdog chan os.Signal, err *error) {
    http.HandleFunc("/", eutherpeHTTPHandler.handler)
    (*err) = http.ListenAndServe(eutherpeHTTPHandler.eutherpeVars.HTTPd.Addr, nil)
    // TODO(Rafael): Setup TLS server possibility.
    if (*err) != nil {
        sigintWatchdog <- syscall.SIGINT
    }
}

func (ehh *EutherpeHTTPHandler) handler(w http.ResponseWriter, r *http.Request) {
    var templatedOutput string
    switch r.URL.Path {
        case "/eutherpe":
                var contentType = "text/html"
                if r.Method == "GET" {
                    templatedOutput = ehh.eutherpeVars.HTTPd.IndexHTML
                    if len(ehh.eutherpeVars.CurrentConfig) == 0 {
                        ehh.eutherpeVars.CurrentConfig = vars.EutherpeWebUIConfigSheetDefault
                    }
                } else if r.Method == "POST" {
                    r.ParseForm()
                    actionHandler := actions.GetEutherpeActionHandler(&r.Form)
                    templatedOutput = ehh.processAction(actionHandler, &r.Form)
                    contentType = actions.GetContentTypeByActionId(&r.Form)
                } else {
                    templatedOutput = ehh.eutherpeVars.HTTPd.ErrorHTML
                    ehh.eutherpeVars.LastError = fmt.Errorf("501 Not Implemented")
                }
                w.Header().Set("content-type", contentType)
            break
        default:
            templatedOutput = ehh.processGET(&w, r)
    }
    fmt.Fprintf(w, "%s", renders.RenderData(templatedOutput, ehh.eutherpeVars))
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
