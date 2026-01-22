package theme

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

// Theme represents a site theme
type Theme struct {
	Name          string
	Dir           string
	IndexTemplate *template.Template
	PostTemplate  *template.Template
	PageTemplate  *template.Template
}

// Load loads a theme from the specified directory
func Load(themeDir, themeName string) (*Theme, error) {
	themePath := filepath.Join(themeDir, themeName)

	// Check if theme exists
	if _, err := os.Stat(themePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("theme not found: %s", themePath)
	}

	// Load templates separately to avoid conflicts with {{define "content"}}
	templatesDir := filepath.Join(themePath, "templates")
	basePath := filepath.Join(templatesDir, "base.html")

	// Load index template
	indexTmpl, err := template.ParseFiles(basePath, filepath.Join(templatesDir, "index.html"))
	if err != nil {
		return nil, fmt.Errorf("failed to load index template: %w", err)
	}

	// Load post template
	postTmpl, err := template.ParseFiles(basePath, filepath.Join(templatesDir, "post.html"))
	if err != nil {
		return nil, fmt.Errorf("failed to load post template: %w", err)
	}

	// Load page template
	pageTmpl, err := template.ParseFiles(basePath, filepath.Join(templatesDir, "page.html"))
	if err != nil {
		return nil, fmt.Errorf("failed to load page template: %w", err)
	}

	return &Theme{
		Name:          themeName,
		Dir:           themePath,
		IndexTemplate: indexTmpl,
		PostTemplate:  postTmpl,
		PageTemplate:  pageTmpl,
	}, nil
}

// CopyStatic copies static assets to the output directory
func (t *Theme) CopyStatic(outputDir string) error {
	staticDir := filepath.Join(t.Dir, "static")
	outputStaticDir := filepath.Join(outputDir, "static")

	// Check if static directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		return nil // No static files, not an error
	}

	// Create output static directory
	if err := os.MkdirAll(outputStaticDir, 0755); err != nil {
		return err
	}

	// Walk through static directory and copy files
	return filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Calculate relative path
		relPath, err := filepath.Rel(staticDir, path)
		if err != nil {
			return err
		}

		// Create destination path
		destPath := filepath.Join(outputStaticDir, relPath)

		// Create destination directory if needed
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}

		// Copy file
		return copyFile(path, destPath)
	})
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
