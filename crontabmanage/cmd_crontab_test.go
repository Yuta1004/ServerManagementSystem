package crontabmanage

import (
	"reflect"
	"testing"
)

func TestCrontabResultParse(t *testing.T) {
	crontabResult :=
		"0 0,12 */1 * * command\n" +
			"0 0 * * 0 command\n" +
			"0 * * * * command\n" +
			"30 6,11,17 * * * command"
	result := crontabResultParse([]byte(crontabResult))

	if len(result) != 4 {
		t.Fail()
	}
	if result[0].Minute.Wildcard {
		t.Fail()
	}
	if !reflect.DeepEqual(result[0].Minute.Nums, []int{0}) {
		t.Fail()
	}
	if !reflect.DeepEqual(result[0].Hour.Nums, []int{0, 12}) {
		t.Fail()
	}
	if !reflect.DeepEqual(result[0].Day.Nums, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Fatal(result[0])
	}
}
