package resourcemanage

import (
	"testing"
)

func TestResourceInit(t *testing.T) {
	r := AllocResourceStruct()
	if r.CPU.User != 0 {
		t.Fail()
	}
}

func TestResourceUpdate(t *testing.T) {
	r := AllocResourceStruct()
	r.Update()
}
