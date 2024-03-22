package auth

import (
    "testing"
    "time"
    "fmt"
)

func TestNewAuthWatchdog(t *testing.T) {
    aw := NewAuthWatchdog(3 * time.Second)
    if aw == nil {
        t.Errorf("NewAuthWatchdog() is not returning a valid pointer.\n")
    }
}

func TestAuthWatchdogOnOffDynamics(t *testing.T) {
    aw := NewAuthWatchdog(3 * time.Second)
    aw.On()
    fmt.Println("=== AuthWatchdog was powered on, powering off within 5 secs, wait...")
    time.Sleep(5 * time.Second)
    aw.Off()
    fmt.Println("=== AuthWatchdog on/off seems to be working fine!")
}

func TestAuthWatchdogRefreshAuthWindow(t *testing.T) {
    aw := NewAuthWatchdog(10 * time.Second)
    aw.On()
    aw.RefreshAuthWindow("42.42.42.42")
    time.Sleep(1 * time.Second)
    aw.RefreshAuthWindow("42.42.42.42")
    aw.Off()
}

func TestAuthWatchdogRefreshIsAuthenticated(t *testing.T) {
    aw := NewAuthWatchdog(5 * time.Second)
    aw.On()
    aw.RefreshAuthWindow("42.42.42.42")
    fmt.Println("=== 42.42.42.42 was registered as an authenticated host during 5 secs.")
    time.Sleep(1 * time.Second)
    if !aw.IsAuthenticated("42.42.42.42") {
        t.Errorf("AuthWatchdog.IsAuthenticated() is returning false when it should return true.\n")
    } else {
        fmt.Println("=== Nice, 42.42.42.42 is being reported as authenticated.")
    }
    if aw.IsAuthenticated("42.42.42.1") {
        t.Errorf("AuthWatchdog.IsAuthenticated() is returning true when it should return false.\n")
    } else {
        fmt.Println("=== Nice, 42.42.42.1 is not being reported as authenticated.")
    }
    fmt.Println("=== Now, let's wait 42.42.42.42 expires its authentication window...")
    time.Sleep(10 * time.Second)
    if aw.IsAuthenticated("42.42.42.42") {
        t.Errorf("AuthWatchdog.IsAuthenticated() is returning true when it should return false.\n")
    } else {
        fmt.Println("=== Nice, 42.42.42.42 is not being reported as authenticated from now on.")
    }
    aw.Off()
}