package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/eliuttth-dev/GoTransmute/internal/converter"
)

func main(){
	// Define flags for command-line options
	inputFile := flag.String("input", "", "Input file or directory path")
	outputFile := flag.String("output", "", "Output file or directory path")
	format := flag.String("format", "", "Output format (e.g., json, csv, html)")
	// filter := flag.String("filter", "", "Filter for conversion")
	// configFile := flag.String("config", "", "Custom configuration file path")
	// overwrite := flag.Bool("overwrite", false, "Overwrite existing files")
	// recursive := flag.Bool("recursive", false, "Recursively convert files in directories")
	unitTests := flag.Bool("unittests", false, "Run unit tests")

	flag.Parse()

	if *unitTests{
		runUnitTests()
		return
	}

	// Validate flags and arguments
	if *inputFile == "" {
		log.Fatal("Error: Input file/directory path is required")
	}

	if *outputFile == "" {
		log.Fatal("Error: Output file/directory path is required")
	}

	if *format == "" {
		log.Fatal("Error: Output format is required")
	}

	// Start conversion process
	fmt.Println("Conversion process started...")
	fmt.Println("Input file: %s/n", *inputFile)
	fmt.Println("Output file: %s/n", *outputFile)
}

func runUnitTests(){
	fmt.Println("Running unit tests...")
}
