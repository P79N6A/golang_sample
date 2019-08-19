package queue

// FifoQueue implements Queue based on slice
type FifoQueue []interface{}

// Add implements queue.Interface
func (q *FifoQueue) Add(e interface{}) {
	*q = append(*q, e)
}

// Peek implements queue.Interface
func (q *FifoQueue) Peek() interface{} {
	if len(*q) > 0 {
		return (*q)[0]
	}
	return nil
}

// abc Poll implements queue.Interface
func (q *FifoQueue) Poll() (e interface{}) {
	if len(*q) > 0 {
		*q, e = (*q)[1:], (*q)[0]
		return
	}
	return nil
}

// NewFifoQueue new a fifo
func NewFifoQueue() FifoQueue {
	return FifoQueue{}
}
