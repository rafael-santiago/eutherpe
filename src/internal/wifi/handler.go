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

func SetWPAPassphrase(ESSID, passphrase string, customPath ...string) error {
    out, err := exec.Command(path.Join(getToolPath(customPath...), "wpa_passphrase"), ESSID, passphrase).CombinedOutput()
    if err != nil {
        return err
    }
    sOut := strings.Split(string(out), "\n")
    credentials := sOut[1] + "\n" + sOut[3] + "\n"
    wpaSupplicantConf := `ctrl_interface=/run/wpa_supplicant
fast_reauth=1
#ap_scan=1
#update_config=1
#country=BR
network={
        scan_ssid=0
        proto=WPA
        key_mgmt=WPA-PSK
        auth_alg=OPEN
    {{.CREDENTIALS}}
}
`
    wpaSupplicantConf = strings.Replace(wpaSupplicantConf, "{{.CREDENTIALS}}", credentials, -1)
    return os.WriteFile("/etc/wpa_supplicant/wpa_supplicant.conf", []byte(wpaSupplicantConf), 0777)
}

func Start(ifaceName string, customPath ...string) (*exec.Cmd, error) {
    exec.Command("sudo", path.Join(getToolPath(customPath...), "systemctl"), "stop", "wpa_supplicant").Run()
    procHandle := exec.Command("sudo", path.Join(getToolPath(customPath...), "wpa_supplicant"), "-c", "/etc/wpa_supplicant/wpa_supplicant.conf", "-i", ifaceName)
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

func getToolPath(customPath ...string) string {
    if len(customPath) > 0 {
        return customPath[0]
    }
    return ""
}
