package auth

import (
    "testing"
)

func TestHashKey(t *testing.T) {
    hashData := HashKey("123mudar*")
    if len(hashData) == 0 {
        t.Errorf("HashKey() seems not to be accordingly work.\n")
    }
}

func TestValidate(t *testing.T) {
    hashData := HashKey("123mudar*")
    if Validate("123mudar*", hashData[:4]) {
        t.Errorf("Validate() has succeeded when it should not.\n")
    }
    if !Validate("123mudar*", hashData) {
        t.Errorf("Validate() has not succeeded when it should.\n")
    }
    if Validate("123mudou", hashData) {
        t.Errorf("Validate() has succeeded when it should not.\n")
    }
}
