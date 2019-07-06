package sync

// Semaphore ...
type Semaphore chan int

// Acquire ...
func (s Semaphore) Acquire() {
	s <- 1
}

// Release ...
func (s Semaphore) Release() {
	<-s
}

// New a Semaphore
func New(size int) Semaphore {
	return make(Semaphore, size)
}
