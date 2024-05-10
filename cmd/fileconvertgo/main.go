package main

import (
	"flag"
	"fmt"
	"log"
	"../../internal/converter"
)

func main(){
	// Define flags for command-line options
	inputFile := flag.String("input", "", "Input file or directory path")
	outputFile := flag.String("output", "", "Output file or directory path")
	format := flag.String("format", "", "Output format (e.g., csv=json, json=csv, jsom=md)")
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




	// Convert file from CSV to JSON
  if *format == "csv=json" {
    if err := converter.ConvertCSVToJSON(*inputFile, *outputFile); err != nil {
      log.Fatal("Error converting CSV to JSON", err)
    }
  }
  
  // Convert file from CSV to Markdown
  if *format == "csv=md" {
    if err := converter.ConvertCSVToMarkdown(*inputFile, *outputFile); err != nil {
      log.Fatal("Error converting CSV to Markdown", err)
    }
  }

  // Convert file from CSV to HTML
  if *format == "csv=html" {
    if err := converter.ConvertCSVToHTML(*inputFile, *outputFile); err != nil {
      log.Fatal("Error converting CSV to HTML", err)
    }
  }

  // Convert file from JSON to CSV
  if *format == "json=csv" {
    if err := converter.ConvertJSONToCSV(*inputFile, *outputFile); err != nil {
      log.Fatal("Error converting JSON to CSV")
    }
  }

  // Convert file from JSON to Markdown
  if *format == "json=md" {
    if err := converter.ConvertJSONToMarkdown(*inputFile, *outputFile); err != nil {
      log.Fatal("Error converting CSV to Markdown", err)
    }
  }

  // Convert file from JSON to HTML
  if *format == "json=html" {
    if err := converter.ConvertJSONToHTML(*inputFile, *outputFile); err != nil {
      log.Fatal("Error converting JSON to HTML", err)
    }
  }

	fmt.Println("Conversion successful!")
}

func runUnitTests(){
	fmt.Println("Running unit tests...")
}
