package main

import (
	"code.byted.org/gopkg/env"
	"fmt"
	_ "github.com/sirupsen/logrus"
	"time"
)

func main() {
	var s []string
	fmt.Println(len(s), s, s == nil)

	var s1 = []string{}
	fmt.Println(len(s1), s1, s1 == nil)

	d, err := time.ParseDuration("10s")
	fmt.Println(d.Nanoseconds(), err)

	fmt.Println(env.IDC())
}
