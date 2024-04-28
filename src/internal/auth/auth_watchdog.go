package auth

import (
    "time"
    "sync"
    "strings"
)

type AuthWatchdog struct {
    bestBeforeWindowTime time.Duration
    bestBefore map[string]time.Time
    mtx sync.Mutex
    enabled bool
}

func NewAuthWatchdog(bestBeforeWindowTime time.Duration) *AuthWatchdog {
    aw := &AuthWatchdog{}
    aw.bestBeforeWindowTime = bestBeforeWindowTime
    aw.bestBefore = make(map[string]time.Time)
    return aw
}

func (aw *AuthWatchdog) RefreshAuthWindow(remoteAddr string) {
    aw.mtx.Lock()
    defer aw.mtx.Unlock()
    aw.bestBefore[remoteAddr] = time.Now().Add(aw.bestBeforeWindowTime)
}

func (aw *AuthWatchdog) IsAuthenticated(remoteAddr string) bool {
    aw.mtx.Lock()
    defer aw.mtx.Unlock()
    var addr string
    p := strings.Index(remoteAddr, ":")
    if p > -1 {
        addr = remoteAddr[0:p]
    } else {
        addr = remoteAddr
    }
    _, has := aw.bestBefore[addr]
    return has
}

func (aw *AuthWatchdog) On() {
    aw.mtx.Lock()
    defer aw.mtx.Unlock()
    aw.enabled = true
    go func() {
        var shouldExit = false
        for !shouldExit {
            time.Sleep(1 * time.Second)
            aw.mtx.Lock()
            now := time.Now()
            remoteAddrs := make([]string, 0)
            for remoteAddr, bestBefore := range aw.bestBefore {
                if now.After(bestBefore) {
                    remoteAddrs = append(remoteAddrs, remoteAddr)
                }
            }
            for _, remoteAddr := range remoteAddrs {
                delete(aw.bestBefore, remoteAddr)
            }
            shouldExit = !aw.enabled
            aw.mtx.Unlock()
        }
    }()
}

func (aw *AuthWatchdog) Off() {
    aw.mtx.Lock()
    defer aw.mtx.Unlock()
    aw.enabled = false
}
