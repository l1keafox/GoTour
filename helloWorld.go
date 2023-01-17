package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	strArray := strings.Fields(s)
	fmt.Println(strArray)

	for _, v := range strArray {

		_, ok := m[v]
		if ok == false {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
