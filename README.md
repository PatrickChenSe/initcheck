# initcheck
Go linter to check initialization of various Go types, such as slice, map, struct etc,.

## Code Style
initchecker aims to detect "smelly" Go code in initialization. To follow the guidelines from [Uber style guide](https://github.com/uber-go/guide/blob/master/style.md), initchecker detects below code as "bad".
* [nil is a valid slice](https://github.com/uber-go/guide/blob/master/style.md#nil-is-a-valid-slice)
* [Omit Zero Value Fields in Structs](https://github.com/uber-go/guide/blob/master/style.md#omit-zero-value-fields-in-structs)
* [Use var for Zero Value Structs](https://github.com/uber-go/guide/blob/master/style.md#use-var-for-zero-value-structs)
* [Initializing Maps with make](https://github.com/uber-go/guide/blob/master/style.md#initializing-maps)
* [Local Variable Declarations](https://github.com/uber-go/guide/blob/master/style.md#local-variable-declarations)

### nil is a valid slice
```go
// "bad" code
a := []int{}
// "good" code
var a []int
```
### Local Variable Declaration
```go
// "bad" code
var c = "foo"
// "good" code
c := "foo"
```
### Omit Zero Value Fields in Structs
```go
// "bad" code
Order struct {
    id int
    name string
    num int
    price double 
}
o1 := new Order{
    id: 111111,
    name: "",
    num: 0,
    price: 0.0,
}

// "good" code
o1 := new Order{
    id: 111111,
}

```

### Use var for Zero Value Structs
```go
// "bad" code
o2 := Order{}
// "good" code
var o2 Order
```

### Initializing Maps with make
```go
// "bad" code
m1 = map[T1]T2{}
// "good" code
m1 = make(map[T1]T2)
```


## Run initchecker
1. Download the package
```bash
go install github.com/patrickchense/initcheck/cmd@latest
```
2. Run the checker
```bash

```

### Run from source
1. Build the project
```bash
go build -o ./bin/initchecker ./cmd
```
2. Run against a file or package
```bash
./bin/initchecker ./test
```

### Run with Golangci-lint

