package parser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ContentData represents the parsed data from a content file
type ContentData struct {
	// Source file path
	SourcePath string
	// Content type (tuvung, nguphap, etc.)
	ContentType string
	// ID of the content (from filename)
	ContentID string
	// Headers from the file (column names)
	Headers []string
	// Rows of data
	Rows []map[string]string
	// Raw data for custom processing
	RawData interface{}
}

// ParseFile parses a file based on its extension and content type
func ParseFile(filePath, contentType string) (*ContentData, error) {
	ext := strings.ToLower(filepath.Ext(filePath))

	// Create content data with basic info
	data := &ContentData{
		SourcePath:  filePath,
		ContentType: contentType,
		ContentID:   getContentID(filePath),
	}

	switch ext {
	case ".csv":
		return parseCSV(filePath, data)
	case ".json":
		return parseJSON(filePath, data)
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}
}

// getContentID extracts content ID from file name
func getContentID(filePath string) string {
	baseName := filepath.Base(filePath)
	return strings.TrimSuffix(baseName, filepath.Ext(baseName))
}

// parseCSV parses CSV files
func parseCSV(filePath string, data *ContentData) (*ContentData, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)
	
	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("empty CSV file")
	}

	// First row contains headers
	data.Headers = records[0]

	// Process each data row
	for i := 1; i < len(records); i++ {
		row := make(map[string]string)
		for j, value := range records[i] {
			if j < len(data.Headers) {
				row[data.Headers[j]] = value
			}
		}
		data.Rows = append(data.Rows, row)
	}

	return data, nil
}

// parseJSON parses JSON files
func parseJSON(filePath string, data *ContentData) (*ContentData, error) {
	// Read the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse JSON data
	var jsonData interface{}
	if err := json.Unmarshal(fileContent, &jsonData); err != nil {
		return nil, err
	}

	// Store the raw data
	data.RawData = jsonData

	// Handle different JSON structures based on content type
	switch jsonMap := jsonData.(type) {
	case map[string]interface{}:
		// Extract headers and rows based on the JSON structure
		if items, ok := jsonMap["items"].([]interface{}); ok {
			// Try to extract headers from the first item
			if len(items) > 0 {
				if firstItem, ok := items[0].(map[string]interface{}); ok {
					for key := range firstItem {
						data.Headers = append(data.Headers, key)
					}
				}
			}

			// Process each item
			for _, item := range items {
				if itemMap, ok := item.(map[string]interface{}); ok {
					row := make(map[string]string)
					for key, value := range itemMap {
						// Convert value to string
						row[key] = fmt.Sprintf("%v", value)
					}
					data.Rows = append(data.Rows, row)
				}
			}
		}
	case []interface{}:
		// Handle JSON arrays
		if len(jsonMap) > 0 {
			if firstItem, ok := jsonMap[0].(map[string]interface{}); ok {
				for key := range firstItem {
					data.Headers = append(data.Headers, key)
				}
			}

			for _, item := range jsonMap {
				if itemMap, ok := item.(map[string]interface{}); ok {
					row := make(map[string]string)
					for key, value := range itemMap {
						row[key] = fmt.Sprintf("%v", value)
					}
					data.Rows = append(data.Rows, row)
				}
			}
		}
	}

	return data, nil
}

// Parser interface for pluggable parsers
type Parser interface {
	Parse(filePath string) (*ContentData, error)
	SupportedExtensions() []string
}

// RegisterParser registers a new parser
func RegisterParser(p Parser) {
	// This is a placeholder for a plugin system
	// In a future version, parsers could be registered here
} 