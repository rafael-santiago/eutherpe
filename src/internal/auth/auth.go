//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package auth

import (
    "os/exec"
    "math/rand"
    "time"
    "strings"
)

const (
    kAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    kAlphaLen = len(kAlpha)
)

func HashKey(password string) string {
    return hashKey(password, getSalt())
}

func Validate(password, hashData string) bool {
    hashDataFields := strings.Split(hashData, "$")
    if len(hashDataFields) < 3 {
        return false
    }
    reHashData := hashKey(password, hashDataFields[2])
    return (reHashData == hashData)
}

func hashKey(password, salt string) string {
    cmd := exec.Command("openssl", "passwd", "-6", "-salt", salt, "-stdin")
    cmd.Stdin = strings.NewReader(password)
    hashData, err := cmd.CombinedOutput()
    if err != nil {
        return ""
    }
    return strings.Replace(string(hashData), "\n", "", -1)
}

func getSalt() string {
    rand.Seed(time.Now().UnixNano())
    var salt string
    for saltSize := rand.Intn(kAlphaLen); saltSize > 0; saltSize-- {
        salt += string(kAlpha[rand.Intn(kAlphaLen)])
    }
    return salt
}
