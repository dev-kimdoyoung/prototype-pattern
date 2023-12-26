package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 30; i++ {
		singleton := GetInstance()
		fmt.Println(singleton)
	}
}
