package node

type Node struct {
	Name            string
	Ip              string
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	TaskCount       int
}
