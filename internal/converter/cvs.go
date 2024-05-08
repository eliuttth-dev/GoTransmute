package converter

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

func ConvertCSVToJSON(inputFile, outputFile string) error {

	// Open CVS File 
	csvFile, err := os.Open(inputFile)

	if err != nil{
		return err
	}

	defer csvFile.Close()

	// Read the file
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()

	if err != nil {
		return err
	}

	// convert to JSON
	var jsonData []map[string]string
	
	headers := records[0]

	for _, record := range records[1:] {
		
		entry := make(map[string]string)

		for i, value := range record {
			entry[headers[i]] = value
		}
		jsonData = append(jsonData, entry)
	}

	jsonFile, err := os.Create(outputFile)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(jsonData); err != nil {
		return err
	}

	return nil
}

