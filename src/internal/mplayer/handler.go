package mplayer

import (
    "os/exec"
    "path"
)

func Play(filePath string, customPath ...string) (*exec.Cmd, error) {
    var mpg123Path string = "mpg123"
    if len(customPath) > 0 {
        mpg123Path = path.Join(customPath[0], mpg123Path)
    }
    cmd := exec.Command(mpg123Path, filePath)//, "-b", "65535", "--no-seekbuffer", "-y")
    return cmd, cmd.Start()
}

func Stop(handle *exec.Cmd) error {
    return handle.Process.Kill()
}
