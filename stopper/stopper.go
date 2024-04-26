package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to control the timer
	timerControl := make(chan bool, 1) // Adding buffer size 1

	// Start a goroutine to handle the timer
	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				fmt.Println("Timer expired")
			case restart := <-timerControl:
				if restart {
					fmt.Println("Timer restarted")
					timer.Reset(5 * time.Second)
				} else {
					fmt.Println("Timer paused")
				}
			}
		}
	}()

	// Wait for 2 seconds
	time.Sleep(2 * time.Second)

	// Send a signal to pause the timer
	timerControl <- false

	// Wait for another 3 seconds
	time.Sleep(3 * time.Second)

	// Send a signal to restart the timer
	timerControl <- true

	// Wait for the timer to expire again
	time.Sleep(5 * time.Second)
}
