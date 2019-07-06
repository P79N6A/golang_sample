package hello

// IntHeap implements a min-heap of int
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Push implements heap.Interface
// 因为要修改h的值, 所以receiver是指针, 排序接口不需要是因为[]int本身已经是引用了
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop implements heap.Interface
func (h *IntHeap) Pop() (x interface{}) {
	*h, x = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
