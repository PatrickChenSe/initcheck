package test

import "fmt"

func test() {
	ns := []int{}
	ns = append(ns, 1)
	fmt.Println(ns)

	m := map[string]string{}
	m["hello"] = "world"
	fmt.Println(m)

	var s = "foo"
	fmt.Println(s)

	var i = 10
	fmt.Println(i)

	b := true
	fmt.Println(b)
}
