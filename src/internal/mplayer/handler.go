package mplayer

import (
    "os/exec"
    "path"
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
