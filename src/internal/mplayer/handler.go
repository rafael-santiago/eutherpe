package mplayer

import (
    "os/exec"
    "path"
    "strconv"
    "strings"
)

func Play(filePath string, customPath ...string) (*exec.Cmd, error) {
    var ffplayPath string = "ffplay"
    if len(customPath) > 0 {
        ffplayPath = path.Join(customPath[0], ffplayPath)
    }
    cmd := exec.Command(ffplayPath, filePath, "-nodisp", "-autoexit")
    return cmd, cmd.Start()
}

func Stop(handle *exec.Cmd) error {
    return handle.Process.Kill()
}

func SetVolume(percentage int, customPath ...string) {
    var pactlPath string = "pactl"
    if len(customPath) > 0 {
        pactlPath = path.Join(customPath[0], pactlPath)
    }
    exec.Command(pactlPath, "--", "set-sink-volume", "0", "-100%").Run()
    exec.Command(pactlPath, "--", "set-sink-volume", "0", "+" + strconv.Itoa(percentage) + "%").Run()
}

func GetVolumeLevel(customPath ...string) uint {
    var amixerPath string = "amixer"
    if len(customPath) > 0 {
        amixerPath = path.Join(customPath[0], amixerPath)
    }
    out, err := exec.Command(amixerPath).CombinedOutput()
    if err != nil {
        return 0
    }
    sOut := string(out)
    off := strings.Index(sOut, "Front Left: Playback")
    if off == -1 {
        return 0
    }
    var vol string
    for ; off < len(sOut); off++ {
        if sOut[off] == '[' {
            off += 1
            for ; off < len(sOut) && sOut[off] >= '0' && sOut[off] <= '9'; off++ {
                vol += string(sOut[off])
            }
            break
        }
    }
    v, err := strconv.Atoi(vol)
    if err != nil {
        v = 0
    }
    return uint(v)
}
