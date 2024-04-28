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