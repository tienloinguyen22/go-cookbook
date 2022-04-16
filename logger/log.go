package logger

import (
	"bytes"
	"fmt"
	"log"
)

func myLogger() {
	buf := bytes.Buffer{}
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)
	logger.Println("test")
	logger.SetPrefix("new logger: ")
	logger.Printf("you can add args too (%v) and use Fatalln to log & crash", true)
	fmt.Println(buf.String())
}