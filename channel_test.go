package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello Muhammad Syaugi"
		fmt.Println("selesai menjalankan channel")

	}()

	fmt.Println("Ini channel nya ", <-channel)

	fmt.Println("mantab")
	time.Sleep(5 * time.Second)

}

/* channel sebagai parameter */

func GiveMeResponse(channelParam chan string) {
	time.Sleep(2 * time.Second)

	channelParam <- "Hello Muhammad Syaugi"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)

	fmt.Println("selesai menampilkan data channel")

	time.Sleep(5 * time.Second)

	fmt.Println("selesai function")

}

/* ./ channel sebagai parameter */

/* channel in dan out pada parameter */

// untuk menerima param channel sebagai channel yang menerima data
func onlyIn(channelParam chan<- string) {
	time.Sleep(2 * time.Second)
	channelParam <- "Hello ini, ini data channel in"
	fmt.Println("selesai menambah data ke channel")
}

func onlyOut(channelParam <-chan string) {
	data := <-channelParam
	fmt.Println(data)
	fmt.Println("selesai menampilkan data dari channel")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go onlyIn(channel)
	go onlyOut(channel)

	fmt.Println("sebelum sleep 5 detik")

	time.Sleep(5 * time.Second)

	fmt.Println("selesai function")

}

/* ./ channel in dan out pada parameter */
