package utils

import (
	"log"
	"os"
	"strconv"
	"time"
)

// timeTrack Print a log with time start and end of function
func timeTrack(start time.Time, name string) {

	servicelog, _ := strconv.ParseBool(os.Getenv("LOG"))
	if servicelog {
		elapsed := time.Since(start)
		log.Printf("[SERVICE] %s start at %s and took %s", name, start, elapsed)
	}

}
