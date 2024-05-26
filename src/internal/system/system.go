package system

import (
    "os/exec"
    "path"
    "flag"
)

func Shutdown() error {
    return doShutdown("-h", "now")
}

func Reboot() error {
    return doShutdown("-r", "now")
}

func doShutdown(params ...string) error {
    var customPath string
    var sudo string
    if flag.Lookup("test.v") != nil {
        customPath = "../system"
    } else {
        sudo = "sudo "
    }
    return exec.Command(sudo + path.Join(customPath, "shutdown"), params...).Run()
}
