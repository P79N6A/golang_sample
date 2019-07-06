package tree

func ExampleTree() {
	// tree
	tr := &Tree{
		Less: func(x, y interface{}) bool {
			v1, ok1 := x.(int)
			v2, ok2 := y.(int)
			if ok1 && ok2 {
				return v1 < v2
			}
			return false
		},
	}
	values := []int{4, 1, 6, 2, 8, 9, 3}
	for _, v := range values {
		tr.Add(v, v)
	}

	tr.Traverse()
	tr.TraverseByLevel()
}
