package queue

// Interface of queue
type Interface interface {
	// Add e to the head of this queue
	Add(e interface{})
	// Peek return the head of this queue, or nil if this queue is empty
	Peek() interface{}
	// Poll return and remove the head of this queue, or nil if this queue is empty
	Poll() interface{}
}
