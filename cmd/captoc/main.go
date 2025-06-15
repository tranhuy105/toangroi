package main

import (
	"fmt"
	"os"
	"path/filepath"

	"captoc/internal/config"
	"captoc/internal/parser"
	"captoc/internal/template"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "build":
		build()
	case "preview":
		preview()
	case "clean":
		clean()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: captoc <command>")
	fmt.Println("Commands:")
	fmt.Println("  build    Generate static website from data files")
	fmt.Println("  preview  Start a local server to preview the website")
	fmt.Println("  clean    Remove generated output files")
}

func build() {
	fmt.Println("Building static website...")

	// Load configuration
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuration loaded successfully.")

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Output directory created.")

	// Copy static assets
	if err := copyStaticAssets(cfg); err != nil {
		fmt.Printf("Error copying static assets: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Static assets copied.")

	// Process data files
	fmt.Println("Processing data files...")
	if err := processDataFiles(cfg); err != nil {
		fmt.Printf("Error processing data files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Data files processed.")

	// Generate index page
	fmt.Println("Generating index page...")
	if err := generateIndexPage(cfg); err != nil {
		fmt.Printf("Error generating index page: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Build completed successfully!")
}

func preview() {
	// Load configuration
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting local preview server...")
	// Check if output directory exists
	if _, err := os.Stat(cfg.OutputDir); os.IsNotExist(err) {
		fmt.Println("Output directory doesn't exist. Running build first...")
		build()
	}

	// Start a simple HTTP server
	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")
	err = startServer(cfg.OutputDir, 8080)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

func clean() {
	// Load configuration
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Cleaning output directory...")
	if err := os.RemoveAll(cfg.OutputDir); err != nil {
		fmt.Printf("Error cleaning output directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Output directory cleaned successfully!")
}

func copyStaticAssets(cfg *config.Config) error {
	// Create static directory in output
	staticOutputDir := filepath.Join(cfg.OutputDir, "static")
	if err := os.MkdirAll(staticOutputDir, 0755); err != nil {
		return err
	}

	// Create subdirectories for CSS and JS
	cssOutputDir := filepath.Join(staticOutputDir, "css")
	jsOutputDir := filepath.Join(staticOutputDir, "js")
	
	if err := os.MkdirAll(cssOutputDir, 0755); err != nil {
		return err
	}
	
	if err := os.MkdirAll(jsOutputDir, 0755); err != nil {
		return err
	}

	// Copy CSS, JS files from templates/static to output/static
	staticSourceDir := filepath.Join("templates", "static")
	
	// Handle copying with explicit errors
	err := filepath.Walk(staticSourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %w", path, err)
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read source file
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", path, err)
		}

		// Get relative path and construct destination path
		relPath, err := filepath.Rel(staticSourceDir, path)
		if err != nil {
			return fmt.Errorf("error computing relative path for %s: %w", path, err)
		}
		
		destPath := filepath.Join(staticOutputDir, relPath)
		
		// Ensure destination directory exists
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", destDir, err)
		}

		// Write file to destination
		if err := os.WriteFile(destPath, data, 0644); err != nil {
			return fmt.Errorf("error writing file %s: %w", destPath, err)
		}
		
		fmt.Printf("  Copied static file: %s -> %s\n", path, destPath)
		
		return nil
	})
	
	if err != nil {
		return fmt.Errorf("error copying static assets: %w", err)
	}
	
	return nil
}

func processDataFiles(cfg *config.Config) error {
	// Process all data directories
	dataDirs, err := os.ReadDir(cfg.DataDir)
	if err != nil {
		return err
	}

	for _, dir := range dataDirs {
		if !dir.IsDir() {
			continue
		}

		dirPath := filepath.Join(cfg.DataDir, dir.Name())
		files, err := os.ReadDir(dirPath)
		if err != nil {
			return err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			// Skip non-CSV files for now
			if filepath.Ext(file.Name()) != ".csv" {
				continue
			}

			filePath := filepath.Join(dirPath, file.Name())
			contentType := dir.Name() // "tuvung" or "nguphap"

			// Parse the file
			fmt.Printf("  Parsing file: %s\n", filePath)
			data, err := parser.ParseFile(filePath, contentType)
			if err != nil {
				fmt.Printf("Warning: Error parsing %s: %v\n", filePath, err)
				continue
			}

			// Generate HTML from template
			outputPath := filepath.Join(
				cfg.OutputDir,
				contentType,
				fmt.Sprintf("%s.html", filepath.Base(file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))])),
			)

			// Create output directory if it doesn't exist
			fmt.Printf("  Creating directory: %s\n", filepath.Dir(outputPath))
			if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(outputPath), err)
			}

			// Render template
			fmt.Printf("  Rendering template for: %s\n", filePath)
			if err := template.RenderTemplate(contentType, data, outputPath, cfg); err != nil {
				return fmt.Errorf("failed to render template for %s: %w", filePath, err)
			}
			fmt.Printf("  Successfully rendered: %s\n", outputPath)
		}
	}

	return nil
}

func generateIndexPage(cfg *config.Config) error {
	// Generate index page with links to all content
	return template.RenderIndex(cfg)
}

// startServer is implemented in server.go 