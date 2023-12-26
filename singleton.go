package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct {
}

var singleInstance *singleton

func GetInstance() *singleton {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &singleton{}
		})
	} else {
		fmt.Println("single instance is already created")
	}

	return singleInstance
}
