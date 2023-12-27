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

func GiveMeResponse(channel chan string) { // channel will be "pass by reference" as a default, so we don't need to use pointer on using channel as parameter
	time.Sleep(2 * time.Second)
	channel <- "Jane Doe"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
