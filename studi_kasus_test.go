package belajar_golang_goroutine

// import (
// 	"fmt"
// 	"strconv"
// 	"testing"
// 	"time"
// )

// func ADisplay() {
// 	channel := make(chan string)

// 	func() {
// 		for i_a := 0; i_a < 5; i_a++ {
// 			channel <- "AAA data ke : " + strconv.Itoa(i_a)
// 		}

// 		close(channel)
// 	}()

// 	index := 0
// 	for datas := range channel {
// 		fmt.Print("AAA index ke :", index, " :")
// 		fmt.Println("AAA menerima data", datas)
// 		index++
// 	}

// }

// func BDisplay() {
// 	channel := make(chan string)

// 	func() {
// 		for i_b := 0; i_b < 5; i_b++ {
// 			channel <- "BBB data ke : " + strconv.Itoa(i_b)
// 		}

// 		close(channel)
// 	}()

// 	index := 0
// 	for datas := range channel {
// 		fmt.Print("BBB index ke :", index, " :")
// 		fmt.Println("BBB menerima data", datas)
// 		index++
// 	}
// }

// func TestCobaRangeChannel(t *testing.T) {

// 	go ADisplay()
// 	go BDisplay()

// 	time.Sleep(2 * time.Second)
// 	fmt.Println("selesai")

// }

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// Struktur data untuk Pesanan
type Order struct {
	ID        int
	Name      string
	Detail    string
	StartTime time.Time // Waktu mulai pemrosesan pesanan
	EndTime   time.Time // Waktu selesai pemrosesan pesanan
	WorkerID  int       // ID dari worker yang memproses pesanan
}

// Fungsi worker untuk memproses pesanan
func worker(id int, orders <-chan Order, results chan<- Order) {
	for order := range orders {
		// Simulasi pemrosesan pesanan
		randomNumber := rand.Intn(1000-500+1) + 500 // random milliseconds
		time.Sleep(time.Duration(randomNumber) * time.Millisecond)

		// Tandai waktu selesai pemrosesan dan worker ID
		order.EndTime = time.Now()
		order.WorkerID = id

		results <- order
	}
}

func TestCobaRangeChannel(t *testing.T) {
	// Inisialisasi channel
	orderChannel := make(chan Order, 30)
	resultChannel := make(chan Order, 30)

	// Jumlah worker
	numWorkers := 5

	// Menjalankan worker goroutine
	for i := 1; i <= numWorkers; i++ {
		go worker(i, orderChannel, resultChannel)
	}

	// Mengirimkan pesanan ke channel
	for i := 1; i <= 30; i++ {
		order := Order{
			ID:        i,
			Detail:    fmt.Sprintf("Detail for order %d", i),
			StartTime: time.Now(), // Tandai waktu mulai pemrosesan
		}
		orderChannel <- order
	}

	// Menutup channel pesanan setelah semua pesanan dikirimkan
	close(orderChannel)

	// Mengumpulkan hasil dari channel hasil
	for i := 1; i <= 30; i++ {
		order := <-resultChannel
		duration := order.EndTime.Sub(order.StartTime)
		fmt.Printf("Order ID %d processed by Worker %d in %v\n", order.ID, order.WorkerID, duration)
	}

	// Menutup channel hasil
	close(resultChannel)

	fmt.Println("Selesai memproses semua pesanan.")
}

/* test goroutine */

func loopSeratus() {

	startTime := time.Now()

	var hasilJumlah string

	for i := 0; i < 10000; i++ {
		hasilJumlah = strconv.Itoa(1000000*2 + i)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("seratus", hasilJumlah)
	fmt.Println("durasi : ", duration.Nanoseconds())
	fmt.Println("")
}

func loopDuaRatus() {

	startTime := time.Now()

	var hasilJumlah string

	for i := 0; i < 29000; i++ {
		hasilJumlah = strconv.Itoa(2000000*2 + i)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("dua ratus", hasilJumlah)
	fmt.Println("durasi : ", duration.Nanoseconds())
	fmt.Println("")
}

func loopTigaRatus() {

	startTime := time.Now()

	var hasilJumlah string

	for i := 0; i < 3000; i++ {
		hasilJumlah = strconv.Itoa(3000000*2 + i)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("tiga ratus", hasilJumlah)
	fmt.Println("durasi : ", duration.Nanoseconds())
	fmt.Println("")
}

func TestGoroutine(t *testing.T) {

	go loopSeratus()
	go loopDuaRatus()
	go loopTigaRatus()

	time.Sleep(5 * time.Second)
}

/* ./ test goroutine */

/* test durasi */

func TestDurasi(t *testing.T) {
	startTime := time.Now()

	for i := 0; i < 400000; i++ {
		// fmt.Println("dfasdfsad")
	}

	endTime := time.Now()

	durasi := endTime.Sub(startTime)

	fmt.Println(durasi)
}

/* ./ test durasi */
