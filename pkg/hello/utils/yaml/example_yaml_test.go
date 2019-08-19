package yaml_test

import (
	"code.byted.org/hotsoon/monit_aggregator/utils/yaml"
	"fmt"
)

func ExampleYaml() {
	yml, _ := yaml.New("testdata/test.yml")
	if s, ok := yml.GetString("group1.k1"); ok {
		fmt.Println(s)
	}

	if i, ok := yml.GetInt("group2.k2"); ok {
		fmt.Println(i)
	}

	if d, ok := yml.GetTimeDuration("group3.k3"); ok {
		fmt.Println(d.Seconds())
	}

	if b, ok := yml.GetBool("group4.k4"); ok {
		fmt.Println(b)
	}
	// Output:
	// v1
	// 100
	// 10
	// true
}
