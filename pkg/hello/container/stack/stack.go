package stack

// Stack based on slice
type Stack []interface{}

// Push element to stack top
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

// Pop return and remove the top element
func (s *Stack) Pop() (e interface{}, ok bool) {
	if len(*s) > 0 {
		*s, e = (*s)[:len(*s)-1], (*s)[len(*s)-1]
		return e, true
	}
	return nil, false
}

// Peek return the top element
func (s *Stack) Peek() (e interface{}, ok bool) {
	if len(*s) > 0 {
		return (*s)[len(*s)-1], true
	}
	return nil, false
}

// New a Stack with values
func New(value ...interface{}) Stack {
	s := Stack{}
	for _, e := range value {
		s.Push(e)
	}
	return s
}
