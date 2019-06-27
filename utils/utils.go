package utils

import (
	"fmt"
	"io"
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

// GenerateHandler To generate a handler and CRUD to a new model
func GenerateHandler(handler string) {

	// Open handlers.template
	fi, err := os.Open("handlers.template")
	if err != nil {
		panic(err)
	}

	// defer to close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file
	fo, err := os.Create("./handlers/__" + handler + ".go")
	if err != nil {
		panic(err)
	}

	// defer to close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// TEMP
		log.Println(string(buf))

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
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
