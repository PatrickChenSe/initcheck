package d

import "fmt"

func test() {
	var s = "hello" // want "consider using ':=' for variable initialization for conciseness"
	fmt.Println(s)
	// pass without error
	d := 1
	fmt.Println(d)
}
