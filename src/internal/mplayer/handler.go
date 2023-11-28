package mplayer

import (
    "os/exec"
)

func Play(filePath string) (*exec.Cmd, error) {
    cmd := exec.Command("mpg123", filePath)
    return cmd, cmd.Start()
}

func Stop(handle *exec.Cmd) error {
    return handle.Process.Kill()
}
