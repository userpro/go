package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello")
	}()

	time.Sleep(time.Second)
}
