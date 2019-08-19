package resourcemanage

import (
	"testing"
)

func TestVmstatParse(t *testing.T) {
	vmstatResultStr :=
		"procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----\n" +
			" r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st\n" +
			" 2  0 585860   6484  19780 169544    0    0     4     4   11    4  1  0 99  0  0"
	vmstatResult := []byte(vmstatResultStr)
	vmstatInfo := vmstatParse(vmstatResult)

	if vmstatInfo["MemFree"] != 6484 {
		t.Fail()
	}
	if vmstatInfo["CPUUser"] != 1 {
		t.Fail()
	}
	if vmstatInfo["IOIn"] != 4 {
		t.Fail()
	}
}
