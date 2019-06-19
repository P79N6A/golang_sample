package main

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/lxb31/golang-sample/hello"
)

var (
	a = hello.Vertex{}
	b = hello.Vertex{X: 1, Y: 2}
	c = hello.Vertex{Y: 1}
	d = &hello.Vertex{X: 1, Y: 2}
)

//closure
//Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。
//函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	msg := hello.NewMessage("hello")
	//msg.
	r, err := hello.SayHello(msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(&r)
	}

	//Stringer
	v := hello.Vertex{X: 1, Y: 3}
	fmt.Println(v.String())
	fmt.Println(&v) //Print严格遵守接受者的方法集, 而不会进行(&v).Sting()的语法糖扩展

	/*
	* time
	 */
	{
		time := time.Now()
		fmt.Println(time)
		fmt.Println(time.Second())
		fmt.Println(time.Unix())
		fmt.Println(time.UnixNano() / 1000000)
	}

	/*
	* array
	 */
	{
		a1 := [2]string{}
		a2 := [2]string{"hello", "world"}

		a1[0] = "hello"
		a1[1] = "world"
		fmt.Printf("a1=%v, a2=%v, len(a1)=%v, len(a2)=%v\n", a1, a2, len(a1), len(a2))
	}

	/*
	* slice
	 */
	{
		s1 := []int{}
		s2 := []int{1, 2, 3, 4}
		s3 := make([]int, 4) //len=4, cap=4

		s1 = append(s1, 1, 2)
		s3[0] = 1
		s3[1] = 2

		fmt.Printf("s1=%v, s2=%v, s3=%v\n", s1, s2, s3)
		fmt.Printf("len(s1)=%v, len(s2)=%v\n", len(s1), len(s2))
		fmt.Printf("cap(s1)=%v, cap(s2)=%v\n", cap(s1), cap(s2))
		fmt.Printf("s2[2:3]=%v\n", s2[2:3])
	}

	/*
	* range
	 */
	{
		s := []int{1, 2, 3, 4}
		for i, v := range s {
			fmt.Println(i, v)
		}
		for _, v := range s {
			fmt.Println(v)
		}
	}

	/*
	* map
	 */
	{
		m1 := map[string]hello.Vertex{}
		m2 := map[string]hello.Vertex{
			"key1": {X: 10, Y: 11},
			"key2": {X: 30, Y: 40},
		}
		m3 := make(map[string]hello.Vertex)

		m3["key1"] = hello.Vertex{X: 1, Y: 2}
		m3["key2"] = hello.Vertex{X: 3, Y: 4}
		delete(m3, "key1")
		elem, ok := m3["key1"]
		fmt.Println(elem, ok)
		for key, value := range m3 {
			fmt.Println(key, value)
		}
		fmt.Println(m1, m2, m3)
	}

	//function pointer
	{
		max := func(i1 int, i2 int) int {
			if i1 > i2 {
				return i1
			}
			return i2
		}
		imax := max(2, 3)
		fmt.Printf("imax=%v \n", imax)
	}

	//closure
	{
		counter := adder()
		for i := 0; i < 10; i++ {
			fmt.Println(counter(i))
		}
	}

	//type assertion
	{
		var i interface{} = "hello"
		s, ok := i.(string)
		fmt.Println(s, ok)
		f, ok := i.(float64)
		fmt.Println(f, ok)
	}

	//type switch
	{
		var v interface{} = "value"
		switch v.(type) {
		case int:
			fmt.Println("type int")
		case string:
			fmt.Println("type string")
		default:
			fmt.Println("type unknown")
		}

	}

	//error
	{
		err := hello.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	//Reader
	reader := strings.NewReader("hello, world")
	bytes := make([]byte, 8)
	for {
		n, err := reader.Read(bytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("read %v bytes \n", n)
	}

	// channel
	// range/close, select

	// tree
	// tree := tree.New(4, 1, 6, 2, 8, 9, 3)
	// tree.Traverse()
	// tree.TraverseByLevel()

	// sync map
	{
		m := sync.Map{}
		m.Store("key1", "value1")
		m.Store("key2", "value2")
		m.Store(3, 4.0)
		e, ok := m.Load("key1")
		e, ok = m.Load("key1")
		e, ok = m.Load("key1")
		fmt.Println(e, ok)
	}
}
