package main

import (
	"fmt"
	"time"
)

func main() {

	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	ch1 := shirtItem.register(observerFirst)
	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("Observer First notified")
			default: // TODO ctx.Done
				// do nothing
			}
		}
	}()
	ch2 := shirtItem.register(observerSecond)
	go func() {
		for {
			select {
			case <-ch2:
				fmt.Println("Observer First notified")
			default: // TODO ctx.Done
				// do nothing
			}
		}
	}()

	shirtItem.updateAvailability()
	time.Sleep(time.Second * 10)
}
