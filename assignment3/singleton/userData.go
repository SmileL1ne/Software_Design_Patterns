package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type userData struct {
	name string
	age  int
}

var instance *userData

func getIntance() *userData {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance userData...")
			instance = &userData{}
		} else {
			fmt.Println("Bruh, instance already created...")
		}
	} else {
		fmt.Println("Bruh, instance already created...")
	}
	return instance
}
