//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package wifi

import (
    "os/exec"
    "strings"
    "path"
    "os"
    "fmt"
)

const (
    WPASupplicantConfFilePath = "/etc/wpa_supplicant/wpa_supplicant.conf"
)

type WLANPlainCredential struct {
    ESSID string
    Passphrase string
}

func GetIfaces(customPath ...string) []string {
    out, err := exec.Command(path.Join(getToolPath(customPath...), "ip"), "link", "show").CombinedOutput()
    if err != nil {
        return make([]string, 0)
    }
    ifaces := make([]string, 0)
    sOut := string(out)
    for s := strings.Index(sOut, ": wl"); s != -1 && s < len(sOut) ; s = strings.Index(sOut[s + 1:], ": wl") {
        sEnd := strings.Index(sOut[s+1:], ":")
        if sEnd == -1 {
            s += 1
            continue
        }
        ifaces = append(ifaces, sOut[s+2:s+sEnd+1])
        s = s + sEnd + 1
    }
    return ifaces
}

func SetIfaceUp(ifaceName string, customPath ...string) error {
    return exec.Command("sudo", path.Join(getToolPath(customPath...), "ip"), "link", "set", "dev", ifaceName, "up").Run()
}

func SetIfaceDown(ifaceName string, customPath ...string) error {
    return exec.Command("sudo", path.Join(getToolPath(customPath...), "ip"), "link", "set", "dev", ifaceName, "down").Run()
}

func GetWPASupplicantConf(ESSID, passphrase string, customPath ...string) (string, error) {
    var credentials string
    var keyMgmt string
    if len(passphrase) > 0 {
        out, err := exec.Command(path.Join(getToolPath(customPath...), "wpa_passphrase"), ESSID, passphrase).CombinedOutput()
        if err != nil {
            return "", err
        }
        sOut := strings.Split(string(out), "\n")
        credentials = sOut[1] + "\n" + sOut[3] + "\n"
        keyMgmt = "WPA-PSK"
    } else {
        credentials = "ssid=\"" + ESSID + "\""
        keyMgmt = "NONE"
    }
    wpaSupplicantConf := `ctrl_interface=/run/wpa_supplicant
fast_reauth=1
#ap_scan=1
#update_config=1
#country=BR
network={
        scan_ssid=0
        proto=WPA
        key_mgmt={{.KEY-MGMT}}
        auth_alg=OPEN
    {{.CREDENTIALS}}
}
`
    wpaSupplicantConf = strings.Replace(wpaSupplicantConf, "{{.KEY-MGMT}}", keyMgmt, 1)
    wpaSupplicantConf = strings.Replace(wpaSupplicantConf, "{{.CREDENTIALS}}", credentials, 1)
    return wpaSupplicantConf, nil
}

func SetWPAPassphrase(ESSID, passphrase string, customPath ...string) error {
    confData, err := GetWPASupplicantConf(ESSID, passphrase, customPath...)
    if err != nil {
        return err
    }
    return os.WriteFile(WPASupplicantConfFilePath, []byte(confData), 0777)
}

func Start(ifaceName, wpaSupplicantConfFilePath string, customPath ...string) (*exec.Cmd, error) {
    exec.Command("sudo", path.Join(getToolPath(customPath...), "systemctl"), "stop", "wpa_supplicant").Run()
    procHandle := exec.Command("sudo", path.Join(getToolPath(customPath...), "wpa_supplicant"), "-c", wpaSupplicantConfFilePath, "-i", ifaceName)
    return procHandle, procHandle.Start()
}

func Stop(handle *exec.Cmd) error {
    if handle == nil {
        return nil
    }
    return handle.Process.Kill()
}

func LeaseAddr(ifaceName string, customPath... string) (string, error) {
    out, err := exec.Command("sudo", path.Join(getToolPath(customPath...), "dhclient"), "-v", ifaceName).CombinedOutput()
    if err != nil {
        return "", err
    }
    sOut := string(out)
    s := strings.Index(sOut, "bound to ")
    if s == -1 {
        return "", fmt.Errorf("Unable to get a valid ip")
    }
    s += 9
    s_end := strings.Index(sOut[s:], " ")
    if s_end == -1 {
        return "", fmt.Errorf("Unable to get a valid ip")
    }
    addr := sOut[s:s+s_end]
    // INFO(Rafael): In raspbian I observed that after ingressing in AP the routes for multicasting
    //               was not being set up and as a result the mDNS stuff was not going up. It was
    //               failing with setsockopt failure.
    if strings.Index(addr, ":") == -1 {
        exec.Command("sudo", path.Join(getToolPath(customPath...), "ip"), "route", "del", "224.0.0.0/4", "dev", ifaceName).Run()
        exec.Command("sudo", path.Join(getToolPath(customPath...), "ip"), "route", "add", "224.0.0.0/4", "dev", ifaceName).Run()
    } else {
        exec.Command("sudo", path.Join(getToolPath(customPath...), "ip"), "route", "del", "ff02::/120", "dev", ifaceName).Run()
        exec.Command("sudo", path.Join(getToolPath(customPath...), "ip"), "route", "add", "ff02::/120", "dev", ifaceName).Run()
    }
    return addr, nil
}

func ReleaseAddr(ifaceName string, customPath... string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "dhclient"), "-r", ifaceName).Run()
}

func GetPlainWLANCredentials(plainCredentialsFilePath string) ([]WLANPlainCredential, error) {
    blob, err := os.ReadFile(plainCredentialsFilePath)
    plainCredentials := make([]WLANPlainCredential, 0)
    if err != nil {
        return plainCredentials, err
    }
    sBlob := string(blob)
    sBlob = strings.Replace(sBlob, "\r", "", -1)
    lines := strings.Split(sBlob, "\n")
    for _, line := range lines {
        shouldSkip := false
        for _, l := range line {
            if l == ' ' || l == '\t' {
                continue
            } else {
                shouldSkip = (l == '#')
                break
            }
        }
        if shouldSkip {
            continue
        }
        tokOff := len(line) - 1
        for ; tokOff > 0; tokOff-- {
            if line[tokOff] == ' ' || line[tokOff] == '\t' {
                break
            }
        }
        var ESSID string
        var pass string
        if tokOff > 0 && (tokOff+1) < len(line) {
            pass = line[tokOff+1:]
        } else {
            tokOff = len(line)
        }
        ESSID = line[:tokOff]
        if len(ESSID) == 0 {
            continue
        }
        plainCredentials = append(plainCredentials, WLANPlainCredential { ESSID, pass })
    }
    return plainCredentials, nil
}

func getToolPath(customPath ...string) string {
    if len(customPath) > 0 {
        return customPath[0]
    }
    return ""
}
