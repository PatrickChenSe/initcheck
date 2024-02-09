package a

import "fmt"

func test() {
	ns := []int{} // want "consider using 'var' for empty slice initialization to avoid unnecessary memory allocation"
	ns = append(ns, 2)
	fmt.Println(ns)
	var ss []string
	// should pass without error
	ss = append(ss, "hello", "world")
	fmt.Println(ss)
}
