package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// TimeTrack Print a log with time start and end of function
func TimeTrack(start time.Time, name string) {

	servicelog, _ := strconv.ParseBool(os.Getenv("LOG"))
	if servicelog {
		elapsed := time.Since(start)
		log.Printf(" [TRACK] %s start at %s and took %s", name, start.Format("2006/01/02 15:04:05"), elapsed)
	}

}

// WriteHandlers Auxiliate on creation of new REST API's
func WriteHandlers(handlers string) {

	// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go

	f, err := os.Create(handlers + ".go")
	if err != nil {
		fmt.Println(err)
		return
	}

	// close fo on exit and check for its returned error
	defer func() {

		if err := f.Close(); err != nil {
			panic(err)
		}

	}()

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}
