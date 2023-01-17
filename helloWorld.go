package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()
	switch {
	case today.Hour() < 12:
		fmt.Println("Good Morning!")
	case today.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
