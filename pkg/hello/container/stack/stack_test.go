package stack

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// TestMain will run before other test
func TestMain(m *testing.M) {
	fmt.Println("-- setup --")
	os.Exit(m.Run())
}

func TestStack(t *testing.T) {
	src := []int{1, 2, 3}
	dest := []int{}
	expected := []int{3, 2, 1}

	s := New()
	for _, v := range src {
		s.Push(v)
	}

	for len(s) > 0 {
		if v, ok := s.Pop(); ok {
			if v1, ok := v.(int); ok {
				dest = append(dest, v1)
			}
		}
	}
	if len(dest) != len(expected) {
		t.Errorf("stack Pop wrong length, expected: %v, got: %v", len(expected), len(dest))
	}
	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("stack Pop wrong element, expected: %v, got: %v", expected, dest)
	}
}

func BenchmarkStack(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Pop()
		s.Pop()
		s.Pop()
	}
}
