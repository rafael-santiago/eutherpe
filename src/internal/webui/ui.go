package webui

import (
    "internal/vars"
    "internal/actions"
    "internal/renders"
    "fmt"
    "os"
    "path"
    "path/filepath"
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
                if r.Method == "GET" {
                    // TODO(Rafael): Set up a default page.
                    templatedOutput = ehh.eutherpeVars.HTTPd.IndexHTML
                } else if r.Method == "POST" {
                    r.ParseForm()
                    actionHandler := actions.GetEutherpeActionHandler(&r.Form)
                    templatedOutput = ehh.processAction(actionHandler, &r.Form)
                } else {
                    templatedOutput = ehh.eutherpeVars.HTTPd.ErrorHTML
                    ehh.eutherpeVars.LastError = fmt.Errorf("501 Not Implemented (ou, boa tentativa mas vai ter que melhorar...)")
                }
                w.Header().Set("content-type", "text/html")
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
                fmt.Errorf("500 Internal Server Error (ou, voce tava muito zureta [sabe-se la de que] quando me mandou isso...)")
        return ehh.eutherpeVars.HTTPd.ErrorHTML
    }
    ehh.eutherpeVars.LastError = actionHandler(ehh.eutherpeVars, userData)
    if ehh.eutherpeVars.LastError != nil {
        userData.Del(vars.EutherpePostFieldLastError)
        userData.Add(vars.EutherpePostFieldLastError, ehh.eutherpeVars.LastError.Error())
    }
    return ehh.eutherpeVars.HTTPd.IndexHTML
}

func (ehh *EutherpeHTTPHandler) processGET(w *http.ResponseWriter, r *http.Request) string {
    vdoc := filepath.Base(r.URL.Path)
    if !ehh.isPubFile(vdoc) {
        ehh.eutherpeVars.LastError = fmt.Errorf("403 Forbidden (ou, sabe de nada inocente...)")
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
