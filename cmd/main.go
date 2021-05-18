package main

import (
	"fmt"
	"DIDWallet/log"
	hublog "github.com/StupidTAO/DIDHub/log"
)

func main() {
	err := log.LogInit()
	if err != nil {
		fmt.Println("panic: log init error")
	}
	err = hublog.LogInit()
	if err != nil {
		fmt.Println("panic: hub log init error")
	}
	Run()
}

