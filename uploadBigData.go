package belajar_golang_goroutine

import (
	"database/sql"
	"fmt"
	"log"
	mathRand "math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/brianvoe/gofakeit/v7"

	_ "github.com/go-sql-driver/mysql"
)

const (
	numOfRecords = 300000
	batchSize    = 10000
)

var x int32 = 0

func generateRandomData() (string, string, string, string, string, int) {

	age := mathRand.Intn(60) + 18

	return gofakeit.Name(),
		gofakeit.Email(),
		gofakeit.StreetName(),
		gofakeit.Phone(),
		gofakeit.Gender(),
		age

}

func insertBatch(db *sql.DB, start, end int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()

	mutex.Lock()

	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return
	}

	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT INTO person (name, email, address, phone, gender, age) VALUES (?, ?, ? ,?, ?, ?)")

	if err != nil {
		log.Println("error prepare statement : ", err)
		return
	}

	defer stmt.Close()

	for i := start; i < end; i++ {
		atomic.AddInt32(&x, 1)

		name, email, address, phone, gender, age := generateRandomData()
		_, err := stmt.Exec(name, email, address, phone, gender, age)

		if err != nil {
			log.Printf("Error executing statement for record %d: %v\n", i, err)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("error Commiting transaction : ", err)
		return
	}

}

func main() {

	// rand.Seed(time.Now().UnixNano())
	var mutex sync.Mutex

	start := time.Now()

	dsn := "root:@tcp(127.0.0.1:3306)/go_dummy"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	var wg sync.WaitGroup

	for i := 0; i < numOfRecords; i += batchSize {
		// for i := 0; i < 10; i++ {
		start := i
		end := i + batchSize
		if end > numOfRecords {
			end = numOfRecords
		}

		wg.Add(1)
		go insertBatch(db, start, end, &wg, &mutex)
	}

	wg.Wait()
	fmt.Println("Data insertion completed.")
	fmt.Println("took : ", time.Since(start))
	fmt.Println("banyak data : ", x)
}

// gofakeit.Name()             // Markus Moen
// gofakeit.Email()            // alaynawuckert@kozey.biz
// gofakeit.Phone()            // (570)245-7485
