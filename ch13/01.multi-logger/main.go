package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	loggingMultipleFilesExample()
	logLevelsExample()
}

func loggingMultipleFilesExample() {
	logFile := new(bytes.Buffer)
	w := SustainedMultiWriter(os.Stdout, logFile)
	l := log.New(w, "example: ", log.Lshortfile|log.Lmsgprefix)

	fmt.Println("standard output:")
	l.Print("Canada is south of Detroit")

	fmt.Print("\nlog file contents:\n", logFile.String())
}

// This example logs DEBUG & ERROR entries to stdout, but only ERROR entries to the log file.
func logLevelsExample() {
	lDebug := log.New(os.Stdout, "DEBUG: ", log.Lshortfile)
	logFile := new(bytes.Buffer)
	w := SustainedMultiWriter(logFile, lDebug.Writer())
	lError := log.New(w, "ERROR: ", log.Lshortfile)

	fmt.Println("standard output:")
	lError.Print("cannot communicate with the database")
	lDebug.Print("you cannot hum while holding your nose")

	fmt.Print("\nlog file contents:\n", logFile.String())
}
