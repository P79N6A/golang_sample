/*
Package designpattern 设计模式
golang 使用组合的方式(嵌入)间接实现继承
嵌入一个结构体以后, 便拥有了该结构体的所有方法与变量, 支持重写
这个多态特性只是golang的语法糖, 结构体初始化时还是需要依次初始化的
golang不支持方法重载
*/
package designpattern

import "fmt"

// Player interface
type Player interface {
	Play()
}

// Person struct
type Person struct {
	Name string
}

// Boy embed Person
type Boy struct {
	Person
}

// Girl embed Person
type Girl struct {
	Person
}

// Woman embed Girl
type Woman struct {
	Girl
}

// Play implementation
func (person *Person) Play() {
	fmt.Printf("Person %v Play\n", person.Name)
}

// Play implementation of Boy
func (boy *Boy) Play() {
	fmt.Printf("Boy %v Play\n", boy.Name)
}

// Play implementation of Girl
func (girl *Girl) Play() {
	fmt.Printf("Girl %v Play\n", girl.Name)
}
