package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// so this method means type T implmenets interface I,
// But we don't need to explicitiy delcare that it does so
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
