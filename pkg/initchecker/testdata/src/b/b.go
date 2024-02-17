package b

import "fmt"

func test() {
	m1 := map[string]int{} // want "consider using 'make' for map initialization to be explicit about intent"
	m1["hello"] = 1
	fmt.Println(m1)
	// should pass without error
	m2 := make(map[string]string)
	m2["hello"] = "world"
	fmt.Println(m2)

}
