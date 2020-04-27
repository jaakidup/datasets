package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func main() {
	PrintMemUsage()

	readStart := time.Now()
	data, err := ioutil.ReadFile("./generated.json")
	if err != nil {
		fmt.Println("Couldn't read file")
	}
	fmt.Println("Time Read", time.Since(readStart))

	unmarshalStart := time.Now()
	var people People
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Couldn't decode json data")
	}
	fmt.Println("Time Unmarshal", time.Since(unmarshalStart))

	count := len(people)
	fmt.Println("People in list", count)

	startRangeLoop := time.Now()
	for _, person := range people {
		person.Name = returner(person.Name)
	}
	fmt.Println("Range Loop time", time.Since(startRangeLoop))

	PrintMemUsage()
}

func returner(input string) string {
	return input
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
