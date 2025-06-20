package template

import (
	"encoding/json"
	"html/template"
	"strings"
	"time"
)

// TemplateFunctions returns a map of custom functions for templates
func TemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"parseOptions": parseOptions,
		"formatYear":   formatYear,
		"capitalize":   capitalize,
		"hasPrefix":    strings.HasPrefix,
		"hasSuffix":    strings.HasSuffix,
		"trimPrefix":   strings.TrimPrefix,
		"trimSuffix":   strings.TrimSuffix,
		"contains":     strings.Contains,
		"split":        strings.Split,
		"slice":        slice,
		"sub":          sub,
		"add":          add,
		"gt":           gt,
		"lt":           lt,
	}
}

// slice returns a substring of s from start to end
func slice(s string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start > end {
		return ""
	}
	return s[start:end]
}

// sub returns x - y
func sub(x, y int) int {
	return x - y
}

// add returns x + y
func add(x, y int) int {
	return x + y
}

// gt returns x > y
func gt(x, y int) bool {
	return x > y
}

// lt returns x < y
func lt(x, y int) bool {
	return x < y
}

// parseOptions parses a string representation of an array into a slice of strings
func parseOptions(s string) []string {
	// Try to parse as JSON first
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		// Replace single quotes with double quotes for JSON parsing
		s = strings.ReplaceAll(s, "'", "\"")
		var options []string
		if err := json.Unmarshal([]byte(s), &options); err == nil {
			return options
		}
	}
	
	// If JSON parsing fails, try simple comma separation
	parts := strings.Split(s, ",")
	options := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			options = append(options, p)
		}
	}
	
	return options
}

// formatYear extracts the year from a date string
func formatYear(date string) string {
	// Try to parse common date formats
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02",
		"2006/01/02",
		"01/02/2006",
	}
	
	for _, f := range formats {
		if t, err := time.Parse(f, date); err == nil {
			return t.Format("2006")
		}
	}
	
	// If parsing fails, return original string
	return date
}

// capitalize capitalizes the first character of a string
func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
} 