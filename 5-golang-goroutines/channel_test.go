package golang_goroutines

import (
	"fmt"
	"strconv"
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

func OnlyIn(channel chan<- string) { // this channel is being used ONLY for IN channel operation
	time.Sleep(2 * time.Second)
	channel <- "John Doe"
}

func OnlyOut(channel <-chan string) { // this channel is being used ONLY for OUT channel operation
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // create buffered channel with "3" as the capacity. this means the channel only able to be sent data 3 times
	defer close(channel)

	channel <- "John"
	channel <- "Jane"
	channel <- "Peter"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("Finished!")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Looping number " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receiving data", data)
	}

	fmt.Println("Finished!")
}
