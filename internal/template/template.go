package template

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"captoc/internal/config"
	"captoc/internal/parser"
)

// TemplateData holds data for template rendering
type TemplateData struct {
	// Configuration
	Config *config.Config
	// Content data
	Content *parser.ContentData
	// Page title
	Title string
	// Current timestamp
	Timestamp string
	// Navigation menu
	Menu []config.MenuItem
	// ContentMap for sidebar navigation
	ContentMap map[string][]string
}

// RenderTemplate renders a template with the given content data
func RenderTemplate(contentType string, data *parser.ContentData, outputPath string, cfg *config.Config) error {
	// Select template file based on content type
	templateFile := filepath.Join(cfg.TemplateDir, contentType+".html")
	goTemplateFile := filepath.Join(cfg.TemplateDir, contentType+".gohtml")
	layoutFile := filepath.Join(cfg.TemplateDir, "layout.gohtml")
	
	// Check if the template file exists, first with .gohtml extension then with .html
	if _, err := os.Stat(goTemplateFile); err == nil {
		templateFile = goTemplateFile
	} else if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		fmt.Printf("Template file %s not found, falling back to default\n", templateFile)
		templateFile = filepath.Join(cfg.TemplateDir, "default.gohtml")
		// If default template doesn't exist either, return an error
		if _, err := os.Stat(templateFile); os.IsNotExist(err) {
			return fmt.Errorf("no template found for content type: %s and default template missing", contentType)
		}
	}
	
	// Check if layout file exists
	if _, err := os.Stat(layoutFile); os.IsNotExist(err) {
		return fmt.Errorf("layout template not found: %s", layoutFile)
	}
	
	fmt.Printf("Using template files: %s and %s\n", layoutFile, templateFile)

	// Find and organize content files for navigation
	contentMap, err := scanContentFiles(cfg)
	if err != nil {
		return err
	}

	// Create template data
	templateData := &TemplateData{
		Config:     cfg,
		Content:    data,
		Title:      fmt.Sprintf("%s - %s", cfg.Name, data.ContentID),
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Menu:       cfg.Menu,
		ContentMap: contentMap,
	}

	// Parse the templates
	tmpl, err := template.New("layout").Funcs(TemplateFunctions()).ParseFiles(layoutFile, templateFile)
	if err != nil {
		return err
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	// Open output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Execute the template
	return tmpl.ExecuteTemplate(outFile, "layout", templateData)
}

// RenderIndex generates the index page
func RenderIndex(cfg *config.Config) error {
	// Load index template
	indexFile := filepath.Join(cfg.TemplateDir, "index.gohtml")
	layoutFile := filepath.Join(cfg.TemplateDir, "layout.gohtml")

	// Check if template exists
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		return fmt.Errorf("index template not found: %s", indexFile)
	}

	// Find and organize content files
	contentMap, err := scanContentFiles(cfg)
	if err != nil {
		return err
	}

	// Create a content object for the index page
	indexContent := &parser.ContentData{
		ContentID:   "index",
		ContentType: "index",
		SourcePath:  "index",
	}

	// Create template data
	templateData := &TemplateData{
		Config:     cfg,
		Content:    indexContent,
		Title:      cfg.Name,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Menu:       cfg.Menu,
		ContentMap: contentMap,
	}

	// Parse the templates
	tmpl, err := template.New("layout").Funcs(TemplateFunctions()).ParseFiles(layoutFile, indexFile)
	if err != nil {
		return err
	}

	// Open output file
	outFile, err := os.Create(filepath.Join(cfg.OutputDir, "index.html"))
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Execute the template
	return tmpl.ExecuteTemplate(outFile, "layout", templateData)
}

// scanContentFiles scans the data directory for content files
func scanContentFiles(cfg *config.Config) (map[string][]string, error) {
	contentMap := make(map[string][]string)

	// Read data directory
	dataDirs, err := os.ReadDir(cfg.DataDir)
	if err != nil {
		return nil, err
	}

	// Process each content type directory
	for _, dir := range dataDirs {
		if !dir.IsDir() {
			continue
		}

		contentType := dir.Name()
		files := []string{}

		// Read content files
		contentFiles, err := os.ReadDir(filepath.Join(cfg.DataDir, contentType))
		if err != nil {
			return nil, err
		}

		for _, file := range contentFiles {
			if file.IsDir() {
				continue
			}

			// Only include supported file types
			ext := filepath.Ext(file.Name())
			if ext == ".csv" || ext == ".json" {
				baseName := file.Name()[:len(file.Name())-len(ext)]
				files = append(files, baseName)
			}
		}

		contentMap[contentType] = files
	}

	return contentMap, nil
}

// RenderString renders a template string with the given data
func RenderString(templateStr string, data interface{}) (string, error) {
	tmpl, err := template.New("inline").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
} 