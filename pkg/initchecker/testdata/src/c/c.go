package c

import "fmt"

type user struct {
	email string
	name  string
}

func test() {
	u := user{} // want "consider using 'var' for zero value struct initialization to improve clarity"
	u.name = "slaise"
	fmt.Println(u.name)
	// pass without error
	var u2 user
	u2.name = "slaise"
	fmt.Println(u2.name)
}
