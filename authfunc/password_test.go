package authfunc

import (
	"testing"
)

func TestAuthPasswordProcess(t *testing.T) {
	password := "GoLanGSeRVerMaNAGe"
	passwordHash := GenPasswordHash(password)
	if !AuthPassword(password, passwordHash) {
		t.Fail()
	}
}