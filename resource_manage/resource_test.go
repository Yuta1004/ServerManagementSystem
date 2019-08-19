package resourcemanage

import (
	"testing"
)

func TestResourceInit(t *testing.T) {
	r := ResourceInit()
	if r.CPU.User != 0 {
		t.Fatal("ResourceInit() test failed!")
	}
}
