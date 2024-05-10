package converter

import (
	"encoding/csv"
	"encoding/json"
	"strings"
  "os"
  "strconv"
  "html/template"
)

func ConvertJSONToCSV(inputFile, outputFile string) error {
	// Open JSON file
	jsonFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Decode JSON
	var jsonData []map[string]interface{}
	err = json.NewDecoder(jsonFile).Decode(&jsonData)
	if err != nil {
		return err
	}

	// Open CSV file
	csvFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write header
	headers := make([]string, 0)
	for key := range jsonData[0] {
		headers = append(headers, key)
	}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Write data
	for _, record := range jsonData {
		row := make([]string, 0)
		for _, header := range headers {
			value, ok := record[header].(string)
			if !ok {
				value = "" // Handle non-string values gracefully
			}
			row = append(row, value)
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func ConvertJSONToMarkdown(inputFile, outputFile string) error {
	// Open JSON file
	jsonFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Decode JSON
	var jsonData []map[string]interface{}
	err = json.NewDecoder(jsonFile).Decode(&jsonData)
	if err != nil {
		return err
	}

	// Open Markdown file
	mdFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer mdFile.Close()

	// Write headers into Markdown file
	headers := make([]string, 0)
	for key := range jsonData[0] {
		headers = append(headers, key)
	}

	mdFile.WriteString("| " + strings.Join(headers, " | ") + " |\n")
	mdFile.WriteString("| " + strings.Repeat("--- |", len(headers)) + " \n")

	// Write rows of data into Markdown file
	for _, record := range jsonData {
		row := make([]string, 0)
		for _, header := range headers {
			// Convert interface{} to string
			var value string
			if strVal, ok := record[header].(string); ok {
				value = strVal
			} else if numVal, ok := record[header].(float64); ok {
				value = strconv.FormatFloat(numVal, 'f', -1, 64)
			} else {
				// Handle other types of data here
				value = ""
			}
			row = append(row, value)
		}
		mdFile.WriteString("| " + strings.Join(row, " | ") + "  |\n")
	}

	return nil
}

func ConvertJSONToHTML(inputFile, outputFile string) error {
  
  // Open JSON file
  jsonFile, err := os.Open(inputFile)

  if err != nil {
    return err
  }

  defer jsonFile.Close()

  // Decode JSON
  var jsonData []map[string]interface{}
  err = json.NewDecoder(jsonFile).Decode(&jsonData)

  if err != nil {
    return err
  }

  // Open HTML file
  htmlFile, err := os.Create(outputFile)

  if err != nil {
    return err
  }

  defer htmlFile.Close()

  htmlTemplate := `
    <!DOCTYPE html>
    <html>
      <head>
        <title>JSON to HTML</title>
        <style>
          table {
            border-collapse: collapse;
            width: 100%;
          }
          th,td {
            border: 1px solid black;
            padding: 8px;
            text-align: left;
          }
          th {
            background-color: #f2f2f2;
          }
        </style>
      </head>
      <body>
        <h1>JSON to HTML</h1>
        <table>
          <tr>
          {{range .Headers}}
            <th>{{.}}</th>
          {{end}}
          </tr>
          {{range .Records}}
          <tr>
          {{range .}}
            <th>{{.}}</th>
          {{end}} 
          </tr>
          {{end}}
        </table>
      </body>
    </html>
  ` 

  // Create a template from HTML string
  tmpl, err := template.New("htmlTemplate").Parse(htmlTemplate)

  if err != nil {
    return err
  }

  // Execute the template with data
  data := struct {
    Headers []string
    Records [][]interface{}    
  } {
    Headers: make([]string, 0),
    Records: make([][]interface{}, 0),
  }

  // Extract headers and records from JSON data
  for _, record := range jsonData {
    row := make([]interface{}, 0)

    for key, value := range record {
      // Extract headers
      if !contains(data.Headers, key) {
        data.Headers = append(data.Headers, key)
      }

      // Extract values
      row = append(row, value)
    }
    data.Records = append(data.Records, row)
  }

  // Execute template
  if err := tmpl.Execute(htmlFile, data); err != nil {
    return err
  }

  return nil
}

// Helper function to checkif a string exists in a slice of strings
func contains(s []string, str string) bool {
  for _, v := range s {
    if v == str {
      return true
    }
  }
  return false
}
