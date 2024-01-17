package storage

import (
    "testing"
    "os"
)

func TestGetAllAvailableStoragesMustFail(t *testing.T) {
    os.Setenv("FINDMNT_MUST_FAIL", "1")
    defer os.Unsetenv("FINDMNT_MUST_FAIL")
    availStorages := GetAllAvailableStorages("./")
    if len(availStorages) != 0 {
        t.Errorf("GetAllAvaialbleStorages() did not fail when it should : '%v'\n", availStorages)
    }
}

func TestGetAllAvailableStoragesMustSucceed(t *testing.T) {
    availStorages := GetAllAvailableStorages("./404")
    if len(availStorages) != 0 {
        t.Errorf("GetAllAvailableStorages() has found what must not!!!! 8|\n")
    }
    availStorages = GetAllAvailableStorages("./mountinfo")
    if len(availStorages) != 2 {
        t.Errorf("GetAllAvaialbleStorages() did not succeed when it should.\n")
    }
    if availStorages[0] != "/" ||
       availStorages[1] != "/media/rs/624B-F629" {
        t.Errorf("GetAllAvaiableStorages() seems not to be returning the exact expected result : '%v'\n", availStorages)
    }
}

