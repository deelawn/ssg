package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/deelawn/ssg/internal/config"
	"github.com/deelawn/ssg/internal/theme"
)

// Generator handles site generation
type Generator struct {
	Config *config.Config
	Theme  *theme.Theme
}

// New creates a new generator
func New(cfg *config.Config, thm *theme.Theme) *Generator {
	return &Generator{
		Config: cfg,
		Theme:  thm,
	}
}

// Build generates the static site
func (g *Generator) Build() error {
	// Clean output directory
	if err := os.RemoveAll(g.Config.Build.OutputDir); err != nil {
		return fmt.Errorf("failed to clean output directory: %w", err)
	}

	// Create output directory
	if err := os.MkdirAll(g.Config.Build.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Copy static assets
	if err := g.Theme.CopyStatic(g.Config.Build.OutputDir); err != nil {
		return fmt.Errorf("failed to copy static assets: %w", err)
	}

	// Process posts
	posts, err := g.processContent("posts", "post.html")
	if err != nil {
		return fmt.Errorf("failed to process posts: %w", err)
	}

	// Process pages
	if _, err := g.processContent("pages", "page.html"); err != nil {
		return fmt.Errorf("failed to process pages: %w", err)
	}

	// Generate index page
	if err := g.generateIndex(posts); err != nil {
		return fmt.Errorf("failed to generate index: %w", err)
	}

	return nil
}

func (g *Generator) processContent(contentType, templateName string) ([]*Content, error) {
	contentDir := filepath.Join(g.Config.Build.ContentDir, contentType)

	// Check if content directory exists
	if _, err := os.Stat(contentDir); os.IsNotExist(err) {
		return nil, nil // No content of this type
	}

	var contents []*Content

	// Walk through content directory
	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-markdown files
		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		// Parse content
		content, err := ParseContent(path)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		// Skip drafts
		if content.Draft {
			fmt.Printf("Skipping draft: %s\n", path)
			return nil
		}

		// Generate URL
		content.GenerateURL(g.Config.Build.ContentDir, contentType)

		// Generate output path
		outputPath := filepath.Join(g.Config.Build.OutputDir, strings.TrimPrefix(content.URL, "/"))

		// Create output directory
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}

		// Render content
		if err := g.renderContent(content, templateName, outputPath); err != nil {
			return fmt.Errorf("failed to render %s: %w", path, err)
		}

		contents = append(contents, content)
		fmt.Printf("Generated: %s\n", outputPath)

		return nil
	})

	return contents, err
}

func (g *Generator) renderContent(content *Content, templateName, outputPath string) error {
	// Create output file
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Prepare template data
	data := map[string]interface{}{
		"Title":       content.Title,
		"Description": content.Description,
		"Date":        content.Date,
		"Author":      content.Author,
		"Content":     content.HTML,
		"Site":        g.Config.Site,
		"Year":        time.Now().Year(),
	}

	// Select the correct template based on content type
	switch templateName {
	case "post.html":
		return g.Theme.PostTemplate.ExecuteTemplate(f, "base.html", data)
	case "page.html":
		return g.Theme.PageTemplate.ExecuteTemplate(f, "base.html", data)
	default:
		return fmt.Errorf("unknown template: %s", templateName)
	}
}

func (g *Generator) generateIndex(posts []*Content) error {
	// Sort posts by date (newest first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	// Create index.html
	outputPath := filepath.Join(g.Config.Build.OutputDir, "index.html")
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Prepare template data
	data := map[string]interface{}{
		"Title": "",
		"Site":  g.Config.Site,
		"Posts": posts,
		"Year":  time.Now().Year(),
	}

	// Execute template
	fmt.Printf("Generated: %s\n", outputPath)
	return g.Theme.IndexTemplate.ExecuteTemplate(f, "base.html", data)
}
