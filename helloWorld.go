package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		45, 44,
	},
	"Google": {
		44, 22,
	},
}

func main() {
	fmt.Println(m)
}
