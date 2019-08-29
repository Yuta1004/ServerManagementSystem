package resource

// Resource : サーバリソースを扱う
type Resource struct {
	CPU    CPUResource
	Memory MemoryResource
	IO     IOResource
	Swap   SwapResource
}

// CPUResource : CPU
type CPUResource struct {
	User   int
	System int
	Idol   int
	Wait   int
}

// MemoryResource : メモリ
type MemoryResource struct {
	Free  int
	Buff  int
	Cache int
}

// IOResource : IO
type IOResource struct {
	In  int
	Out int
}

// SwapResource : スワップ
type SwapResource struct {
	In  int
	Out int
}
