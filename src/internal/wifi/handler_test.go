//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package wifi

import (
    "testing"
    "os"
)

func TestGetIfaces(t *testing.T) {
    os.Setenv("IP_MUST_FAIL", "1")
    ifaces := GetIfaces("../wifi")
    if len(ifaces) != 0 {
        t.Errorf("GetIfaces() has not failed.\n")
    }
    os.Unsetenv("IP_MUST_FAIL")
    ifaces = GetIfaces("../wifi")
    if len(ifaces) != 1 {
        t.Errorf("GetIfaces() has returned a wrong total of interfaces.\n")
    } else if ifaces[0] != "wlxf0a7314a4543" {
        t.Errorf("GetIfaces() has returned an unexpected interface.\n")
    }
}

func TestGetPlainWLANCredentials(t *testing.T) {
    testData := "# ESSID PASSPHRASE\n" +
                "Xablau 1234\n\n\r\n\n\n" +
                "OpenedOne\r\n" +
                "AbcD 42\r\n" +
                "Teste com espacos 424242!"
    err := os.WriteFile("/tmp/pub-aps", []byte(testData), 0777)
    if err != nil {
        t.Errorf("Unable to create /tmp/pub-aps")
    } else {
        defer os.Remove("/tmp/pub-aps")
        credentials, err := GetPlainWLANCredentials("/tmp/pub-aps")
        if err != nil {
            t.Errorf("GetPlainWLANCredentials() returned and error : '%s'.\n", err.Error())
        }
        if len(credentials) != 4 {
            t.Errorf("GetPlainWLANCredentials() returned a wrong total of credentials.\n")
        }
        if credentials[0].ESSID != "Xablau" || credentials[0].Passphrase != "1234" {
            t.Errorf("credentials[0] has unexpected configuration.\n")
        }
        if credentials[1].ESSID != "OpenedOne" || credentials[1].Passphrase != "" {
            t.Errorf("credentials[1] has unexpected configuration.\n")
        }
        if credentials[2].ESSID != "AbcD" || credentials[2].Passphrase != "42" {
            t.Errorf("credentials[2] has unexpected configuration.\n")
        }
        if credentials[3].ESSID != "Teste com espacos" ||
           credentials[3].Passphrase != "424242!" {
            t.Errorf("credentials[3] has unexpected configuration.\n")
        }
    }
}
