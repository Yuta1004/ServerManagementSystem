package resourcemanage

import (
	"fmt"
	"os/exec"
)

// ResourceInit : Resourceを初期化して返す
func ResourceInit() Resource {
	res := Resource{}
	return res
}

// Update : リソース情報を最新のものに更新する
func (r Resource) Update() {
	// vmstat実行
	vmstatResult, err := exec.Command("vmstat").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 実行結果パース
	vmstatInfo, err := vmstatParse(vmstatResult)
	if err != nil {
		fmt.Println(err.Error())
	}

	r.CPU.User = vmstatInfo["CPUUser"]
	r.CPU.System = vmstatInfo["CPUSystem"]
	r.CPU.Idol = vmstatInfo["CPUIdol"]
	r.CPU.Wait = vmstatInfo["CPUWait"]
	r.IO.In = vmstatInfo["IOIn"]
	r.IO.Out = vmstatInfo["IOOut"]
	r.Memory.Buff = vmstatInfo["MemBuff"]
	r.Memory.Cache = vmstatInfo["MemCache"]
	r.Memory.Free = vmstatInfo["MemFree"]
	r.Swap.In = vmstatInfo["SwapIn"]
	r.Swap.Out = vmstatInfo["SwapOut"]
}
