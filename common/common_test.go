package common

import (
	"reflect"
	"testing"
)

func TestDeduplicationArrayInt(t *testing.T) {
	tmp := []int{1, 2, 3, 2, 5, 1, 3, 4, 2, 1, 1, 3, 6, 4}
	tmp = DeduplicationArrayInt(tmp)
	if !reflect.DeepEqual(tmp, []int{1, 2, 3, 4, 5, 6}) {
		t.Fatal(tmp)
	}
}
