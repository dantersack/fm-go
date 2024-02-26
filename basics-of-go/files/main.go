package main

import (
	"log"
	"os"
)

func main() {
	// Get root path
	rootPath, _ := os.Getwd()
	// Read the file from "testdata" folder
	data, readErr := os.ReadFile(rootPath + "/testdata/file.txt")
	if readErr != nil {
		log.Fatal(readErr)
	}
	os.Stdout.Write(data)
	// Create a copy of the file in the same "testdata" folder
	writeErr := os.WriteFile(rootPath+"/testdata/file-copy.txt", data, 0666)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
