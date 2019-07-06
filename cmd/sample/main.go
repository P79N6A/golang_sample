package main

import (
	"fmt"

	_ "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	v1 := pflag.Int("v1", 100, "int value")
	v2 := pflag.StringP("v2", "v", "default", "string value")
	pflag.Parse()
	fmt.Println(*v1, *v2)

}
