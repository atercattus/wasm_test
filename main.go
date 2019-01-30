package main

import (
	"time"
)

func main() {
	go func() {
		for {
			println("Now:", time.Now().Format(time.RFC3339Nano))
			time.Sleep(2 * time.Second)
		}
	}()

	<-JS.waitForUnload()
}
