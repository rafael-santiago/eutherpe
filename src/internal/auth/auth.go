//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package auth

import (
    "math/rand"
    "time"
    "strings"
    "crypto/sha256"
    "encoding/base64"
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
    if len(hashDataFields) < 2 {
        return false
    }
    reHashData := hashKey(password, hashDataFields[1])
    return (reHashData == hashData)
}

func hashKey(password, salt string) string {
    m := len(salt) >> 1
    sum := sha256.Sum256([]byte(salt[:m] + password + salt[m:]))
    strSum := base64.StdEncoding.EncodeToString(sum[:])
    return (strSum + "$" + salt)
}

func getSalt() string {
    rand.Seed(time.Now().UnixNano())
    var salt string
    for saltSize := rand.Intn(kAlphaLen); saltSize > 0; saltSize-- {
        salt += string(kAlpha[rand.Intn(kAlphaLen)])
    }
    return salt
}
