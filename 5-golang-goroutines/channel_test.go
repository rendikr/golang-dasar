package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // no matter what happens, close the channel at the end of this function

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "John Doe" // sending new data to channel
		fmt.Println("Finished sending data to channel")
	}()

	time.Sleep(5 * time.Second)

	data := <-channel // receiving new channel data to a newly var data
	fmt.Println("data:", data)
}
