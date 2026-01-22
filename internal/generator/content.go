package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Content represents a markdown content file
type Content struct {
	Title       string
	Date        string
	Author      string
	Description string
	Draft       bool
	Body        string
	HTML        template.HTML
	URL         string
	FilePath    string
}

// ParseContent parses a markdown file with TOML frontmatter
func ParseContent(filePath string) (*Content, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Split frontmatter and content
	parts := bytes.SplitN(data, []byte("+++"), 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid frontmatter format in %s", filePath)
	}

	// Parse frontmatter
	var frontmatter map[string]interface{}
	if err := toml.Unmarshal(parts[1], &frontmatter); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Get content body
	body := string(bytes.TrimSpace(parts[2]))

	// Create content struct
	content := &Content{
		Body:     body,
		FilePath: filePath,
	}

	// Extract frontmatter fields
	if title, ok := frontmatter["title"].(string); ok {
		content.Title = title
	}
	if date, ok := frontmatter["date"].(string); ok {
		content.Date = date
	}
	if author, ok := frontmatter["author"].(string); ok {
		content.Author = author
	}
	if description, ok := frontmatter["description"].(string); ok {
		content.Description = description
	}
	if draft, ok := frontmatter["draft"].(bool); ok {
		content.Draft = draft
	}

	// Convert markdown to HTML
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allow raw HTML in markdown
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		return nil, fmt.Errorf("failed to convert markdown: %w", err)
	}

	content.HTML = template.HTML(buf.String())

	return content, nil
}

// GenerateURL creates a URL path for the content
func (c *Content) GenerateURL(contentDir, contentType string) {
	rel, _ := filepath.Rel(contentDir, c.FilePath)
	rel = strings.TrimSuffix(rel, filepath.Ext(rel))
	rel = strings.ReplaceAll(rel, "\\", "/")
	c.URL = "/" + rel + ".html"
}
