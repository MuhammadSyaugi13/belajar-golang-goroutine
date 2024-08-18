package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
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

/* buffered channel */

func TestBufferedChannel(t *testing.T) {

	fmt.Println("oke")

	profile := []map[string]interface{}{
		{
			"nama":              "andi fhad9fasulajskca;slkd[as]da[ps[dpausfaiuoskalksjdlakjsdna,sfmas,nfaksifhi3429839a98sa7987(*&*&&^%765)]]",
			"umur":              3647392859320439482,
			"status_perkawinan": true,
		},
		{
			"nama":              "andi",
			"umur":              36,
			"status_perkawinan": true,
		}, {
			"nama":              "andi",
			"umur":              36,
			"status_perkawinan": true,
		},
		{
			"nama":              "andi xhsdjfskdjkjxhkjfhskdjhfkjxhkjhfkdshkfshdkjfhkjhk",
			"umur":              36342352342342,
			"status_perkawinan": true,
		},
	}

	title := "coba buffered channel"

	channel := make(chan interface{}, 2)

	go func() {
		channel <- title
		channel <- profile
		channel <- "hahah"
	}()

	go func() {

		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

/* ./ buffered channel */

/* range channel */

func TestRangeChannel(t *testing.T) {

	channel := make(chan string)

	go func() {
		fmt.Println("jalan go func")
		for i := 0; i < 10000000; i++ {
			fmt.Println("jalan for")
			fmt.Println("")
			channel <- "data ke " + strconv.Itoa(i)
		}

		fmt.Println("close channel")
		fmt.Println("")
		close(channel)
	}()

	index := 1
	for datas := range channel {
		fmt.Print("index ke :", index, " :")
		fmt.Println("menerima data", datas)
		index++
	}

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")

	// for i := 0; i < 10000000; i++ {
	// 	fmt.Println("index ke :", i)
	// }

}

/* ./ range channel */
