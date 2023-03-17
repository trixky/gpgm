package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/trixky/krpsim/checker/checker"
)

func main() {
	// Check if two filenames have been provided as arguments
	if len(os.Args) < 3 {
		fmt.Println("Please provide two filenames as arguments")
		return
	}

	// Read the contents of the first file
	inputFile := os.Args[1]
	inputFileContent, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", inputFile, err)
	}

	// Read the contents of the second file
	outputFile := os.Args[2]
	outputFileContent, err := ioutil.ReadFile(outputFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", outputFile, err)
	}

	// Call the function to process the file contents
	res, err := checker.CheckOutput(string(inputFileContent), string(outputFileContent))
	fmt.Printf("%v %v\n", res, err)
}
