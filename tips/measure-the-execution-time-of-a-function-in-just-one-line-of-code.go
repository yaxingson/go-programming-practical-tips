package main

import "time"
import "fmt"

func TrackTime(startTime time.Time) {
	elapsed := time.Since(startTime)
	fmt.Printf("elapsed: %v\n", elapsed)
}

func main() {
	defer TrackTime(time.Now())

	time.Sleep(500 * time.Millisecond)
}
