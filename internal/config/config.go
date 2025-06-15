package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration
type Config struct {
	// Name of the website
	Name string `yaml:"name"`
	// Base URL for the website (for GitHub Pages or other deployments)
	BaseURL string `yaml:"base_url"`
	// Theme configuration
	Theme ThemeConfig `yaml:"theme"`
	// Navigation menu items
	Menu []MenuItem `yaml:"menu"`
	// Page configurations
	Pages PageConfigs `yaml:"pages"`
	// Directory where data files are stored
	DataDir string `yaml:"data_dir"`
	// Directory where template files are stored
	TemplateDir string `yaml:"template_dir"`
	// Directory where output files will be written
	OutputDir string `yaml:"output_dir"`
}

// ThemeConfig holds the theme configuration
type ThemeConfig struct {
	// Primary color for the theme
	PrimaryColor string `yaml:"primary_color"`
	// Secondary color for the theme
	SecondaryColor string `yaml:"secondary_color"`
	// Text color
	TextColor string `yaml:"text_color"`
	// Muted text color
	TextMuted string `yaml:"text_muted"`
	// Background color
	BackgroundColor string `yaml:"background_color"`
	// Alternative background color
	BackgroundAlt string `yaml:"background_alt"`
	// Border color
	BorderColor string `yaml:"border_color"`
	// Success color
	SuccessColor string `yaml:"success_color"`
	// Error color
	ErrorColor string `yaml:"error_color"`
	// Warning color
	WarnColor string `yaml:"warn_color"`
	// Info color
	InfoColor string `yaml:"info_color"`
	// Font family for the theme
	Font string `yaml:"font"`
}

// MenuItem represents a navigation menu item
type MenuItem struct {
	// Label to display in the menu
	Label string `yaml:"label"`
	// URL to link to
	URL string `yaml:"url"`
	// Icon name
	Icon string `yaml:"icon,omitempty"`
	// Child menu items (submenu)
	Children []MenuItem `yaml:"children,omitempty"`
}

// PageConfigs holds configurations for different page types
type PageConfigs struct {
	// Vocabulary page configuration
	Vocabulary PageConfig `yaml:"vocabulary"`
	// Grammar page configuration
	Grammar PageConfig `yaml:"grammar"`
}

// PageConfig holds configuration for a specific page type
type PageConfig struct {
	// Title of the page
	Title string `yaml:"title"`
	// Whether to show search functionality
	ShowSearch bool `yaml:"show_search"`
	// Card layout type for vocabulary
	CardLayout string `yaml:"card_layout,omitempty"`
	// Whether to show results immediately
	ShowResultImmediately bool `yaml:"show_result_immediately,omitempty"`
	// Whether to highlight correct answers
	HighlightCorrect bool `yaml:"highlight_correct,omitempty"`
	// Field configurations
	Fields []FieldConfig `yaml:"fields"`
}

// FieldConfig holds configuration for a field
type FieldConfig struct {
	// Name of the field
	Name string `yaml:"name"`
	// Display label for the field
	Label string `yaml:"label"`
	// Whether to display the field
	Display bool `yaml:"display"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	cfg := &Config{
		Name:        "CAP TOC Learning Materials",
		BaseURL:     "", // Empty string for local development
		DataDir:     "data",
		TemplateDir: "templates",
		OutputDir:   "output",
	}
	
	// Default theme
	cfg.Theme.PrimaryColor = "#2563eb"
	cfg.Theme.SecondaryColor = "#4b5563"
	cfg.Theme.TextColor = "#111827"
	cfg.Theme.TextMuted = "#6b7280"
	cfg.Theme.BackgroundColor = "#ffffff"
	cfg.Theme.BackgroundAlt = "#f9fafb"
	cfg.Theme.BorderColor = "#e5e7eb"
	cfg.Theme.SuccessColor = "#10b981"
	cfg.Theme.ErrorColor = "#ef4444"
	cfg.Theme.WarnColor = "#f59e0b"
	cfg.Theme.InfoColor = "#3b82f6"
	cfg.Theme.Font = "system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif"

	// Default menu
	cfg.Menu = []MenuItem{
		{
			Label: "Từ vựng",
			URL:   "/tuvung/n1_1.html",
			Icon:  "book",
		},
		{
			Label: "Ngữ pháp",
			URL:   "/nguphap/so1.html",
			Icon:  "pencil",
		},
	}

	// Default vocabulary page config
	cfg.Pages.Vocabulary.Title = "Từ vựng"
	cfg.Pages.Vocabulary.ShowSearch = true
	cfg.Pages.Vocabulary.CardLayout = "flip"
	cfg.Pages.Vocabulary.Fields = []FieldConfig{
		{Name: "japanese", Label: "Kanji", Display: true},
		{Name: "reading", Label: "Reading", Display: true},
		{Name: "meaning", Label: "Meaning", Display: true},
		{Name: "sinoVietnamese", Label: "Hán Việt", Display: true},
	}

	// Default grammar page config
	cfg.Pages.Grammar.Title = "Ngữ pháp"
	cfg.Pages.Grammar.ShowResultImmediately = false
	cfg.Pages.Grammar.HighlightCorrect = true
	cfg.Pages.Grammar.Fields = []FieldConfig{
		{Name: "Câu số", Label: "Question Number", Display: true},
		{Name: "Câu hỏi", Label: "Question", Display: true},
		{Name: "Đáp án đúng", Label: "Correct Answer", Display: false},
		{Name: "Lựa chọn", Label: "Options", Display: true},
	}

	return cfg
}

// Load loads the configuration from a YAML file
func Load(filename string) (*Config, error) {
	// Start with default configuration
	cfg := DefaultConfig()

	// Try to read the configuration file
	data, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		// If the file doesn't exist, create it with default values
		data, err = yaml.Marshal(cfg)
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(filename, data, 0644); err != nil {
			return nil, err
		}
		return cfg, nil
	} else if err != nil {
		return nil, err
	}

	// Parse the YAML data
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	// Set default values for empty fields
	if cfg.DataDir == "" {
		cfg.DataDir = "data"
	}
	if cfg.TemplateDir == "" {
		cfg.TemplateDir = "templates"
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = "output"
	}

	return cfg, nil
}

// Save saves the configuration to a YAML file
func Save(cfg *Config, filename string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
} 