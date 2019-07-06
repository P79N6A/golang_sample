package hello

import "fmt"

// Person ...
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// sort
// -----

// ByAge implements sort.Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// heap
// -----

// PersonPriorityQueue implements heap.Interface
type PersonPriorityQueue []Person

func (q PersonPriorityQueue) Len() int           { return len(q) }
func (q PersonPriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q PersonPriorityQueue) Less(i, j int) bool { return q[i].Age < q[j].Age }

// Push ...
func (q *PersonPriorityQueue) Push(x interface{}) {
	*q = append(*q, x.(Person))
}

// Pop ...
func (q *PersonPriorityQueue) Pop() (p Person) {
	*q, p = (*q)[:len(*q)-1], (*q)[len(*q)-1]
	return
}
