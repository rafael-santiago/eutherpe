//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package mplayer

import (
    "os"
    "os/exec"
    "path"
    "fmt"
    "strconv"
    "strings"
)

func ConvertToMP3(inputPath string, customPath ...string) error {
    var ffmpegPath string = "ffmpeg"
    if len(customPath) > 0 {
        ffmpegPath = path.Join(customPath[0], ffmpegPath)
    }
    ext := path.Ext(inputPath)
    outputPath := strings.Replace(inputPath, ext, ".mp3", -1)
    _, err := os.Stat(outputPath)
    if err == nil {
        // INFO(Rafael): By design we will not re-do a time consuming task if the output
        //               already appears to be there.
        return nil
    }
    return exec.Command(ffmpegPath, "-i", inputPath, outputPath).Run()
}

func Play(filePath string, hasBlueAlsaSink bool, customPath ...string) (*exec.Cmd, error) {
    var mpg123Path string = "mpg123"
    if len(customPath) > 0 {
        mpg123Path = path.Join(customPath[0], mpg123Path)
    }
    var cmd *exec.Cmd
    if hasBlueAlsaSink {
        cmd = exec.Command(mpg123Path, "-o", "alsa:bluealsa", filePath)
    } else {
        cmd = exec.Command(mpg123Path, filePath)
    }
    return cmd, cmd.Start()
}

func Stop(handle *exec.Cmd) error {
    return handle.Process.Kill()
}

func SetVolume(percentage int, mixerControlName string, customPath ...string) {
    var amixerPath string = "amixer"
    if len(customPath) > 0 {
        amixerPath = path.Join(customPath[0], amixerPath)
    }
    sPerc := fmt.Sprintf("%d", percentage)
    if len(mixerControlName) > 0 {
        exec.Command(amixerPath, "-D", "bluealsa", "set", "'" + mixerControlName + "'", sPerc + "%").Run()
    } else {
        exec.Command(amixerPath, "set", "'PCM'", sPerc + "%").Run()
    }
}

func GetVolumeLevel(hasBlueAlsaSink bool, customPath ...string) uint {
    var amixerPath string = "amixer"
    if len(customPath) > 0 {
        amixerPath = path.Join(customPath[0], amixerPath)
    }
    var cmd *exec.Cmd
    if hasBlueAlsaSink {
        cmd = exec.Command(amixerPath, "-D", "bluealsa")
    } else {
        cmd = exec.Command(amixerPath)
    }
    out, err := cmd.CombinedOutput()
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
