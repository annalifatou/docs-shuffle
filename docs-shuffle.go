package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func readCSV(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func contains(arr []int, number int) bool {
	for _, a := range arr {
		if a == number {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Pass a playlist (Example: docs-shuffle.sh python)")
	}

	// Currently just reads from one playlist
	playlist := os.Args[1]

	records := readCSV("playlists/" + playlist + ".csv")
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))

	alreadySeen := []int{}

	for range records {
		index := randomNumber.Int() % len(records)

		// Prevent repeats
		if !contains(alreadySeen, index) {
			fmt.Printf("\r\033[K %s", records[index][0])
		}

		alreadySeen = append(alreadySeen, index)

		time.Sleep(5 * time.Second)
	}
}
