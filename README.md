# CAP TOC - Content Generator

CAP TOC is a flexible content generator for creating static websites from structured data. It's designed to be easily customizable for various content types.

## Features

- Generate static HTML websites from CSV and JSON data files
- Customizable content types with flexible templates
- Mobile-responsive design with light/dark mode support
- Search functionality for content
- Fast local preview server

## Usage

```bash
# Generate the website
captoc build

# Start local preview server
captoc preview

# Clean output files
captoc clean
```

## Content Type System

CAP TOC uses a flexible content type system that allows you to define how your data should be displayed. The system is completely customizable through the `config.yaml` file.

### Creating a New Content Type

1. Create a new directory in the `data` folder with the name of your content type (e.g., `kanji`)
2. Add CSV or JSON files to that directory
3. Add a new entry to the `content_types` section in `config.yaml`

Example configuration:

```yaml
content_types:
  my_content_type:
    title: "My Content"
    template: "default"  # Use the default template or create a custom one
    show_search: true
    fields:
      - name: "field1"
        label: "Field One"
        display: true
      - name: "field2"
        label: "Field Two"
        display: true
```

4. Create a custom template (optional) in the `templates` directory with the name `my_content_type.gohtml`

### Menu Configuration

Add your content type to the navigation menu:

```yaml
menu:
  - label: "My Content"
    url: "/my_content_type/sample.html"
    icon: "icon-name"
    content_type: "my_content_type"
```

## Templates

CAP TOC supports custom templates for each content type. If no template is specified, the default template will be used.

### Available Templates

- `default.gohtml` - A generic template for displaying any content type
- `tuvung.gohtml` - Template optimized for vocabulary display
- `nguphap.gohtml` - Template designed for grammar questions

## CSS Customization

You can customize the appearance by modifying:

- `templates/static/css/style.css` - Main styling
- `templates/static/css/custom.css` - Custom styling for your content types

## License

MIT License 