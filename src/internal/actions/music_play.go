package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "fmt"
    "flag"
    "io/ioutil"
)

func MusicPlay(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../mplayer"
    }
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if eutherpeVars.Player.Handle != nil {
        // INFO(Rafael): Playing already just keep on playing.
        return nil
    }
    upNextLen := len(eutherpeVars.Player.UpNext)
    if upNextLen == 0 {
        return fmt.Errorf("There is no selection to play.")
    }
    if eutherpeVars.Player.UpNextCurrentOffset >= upNextLen {
        eutherpeVars.Player.UpNextCurrentOffset = -1
        eutherpeVars.Player.Stopped = true
        eutherpeVars.Player.NowPlaying = mplayer.SongInfo{}
        return nil
    }
    if eutherpeVars.Player.UpNextCurrentOffset < 0 {
        eutherpeVars.Player.UpNextCurrentOffset = 0
    }
    var err error
    eutherpeVars.Player.NowPlaying = eutherpeVars.Player.UpNext[eutherpeVars.Player.UpNextCurrentOffset]
    createCache(eutherpeVars.Player.NowPlaying.FilePath, "/tmp/cache.mp3")
    eutherpeVars.Player.Handle, err = mplayer.Play("/tmp/cache.mp3"/*eutherpeVars.Player.NowPlaying.FilePath*/, customPath)
    eutherpeVars.Player.Stopped = (err != nil)
    if eutherpeVars.Player.Stopped {
        return err
    }
    go func() {
        if eutherpeVars.Player.Handle == nil {
            return
        }
        eutherpeVars.Player.Handle.Wait()
        if eutherpeVars.Player.Stopped {
            return
        }
        if !eutherpeVars.Player.RepeatOne {
            eutherpeVars.Player.UpNextCurrentOffset++
        }
        if eutherpeVars.Player.RepeatAll &&
           eutherpeVars.Player.UpNextCurrentOffset >= len(eutherpeVars.Player.UpNext) {
            eutherpeVars.Player.UpNextCurrentOffset = -1
        }
        eutherpeVars.Player.Handle = nil
        MusicPlay(eutherpeVars, nil)
    }()
    return nil
}

func createCache(src, dest string) {
    input, _ := ioutil.ReadFile(src)
    ioutil.WriteFile(dest, input, 0644)
}

