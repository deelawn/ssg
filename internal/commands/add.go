package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/deelawn/ssg/internal/config"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add <type> <name>",
	Short: "Add new content to your site",
	Long:  `Create a new post or page with frontmatter template.`,
	Args:  cobra.ExactArgs(2),
	RunE:  runAdd,
}

func runAdd(cmd *cobra.Command, args []string) error {
	contentType := strings.ToLower(args[0])
	name := args[1]

	// Validate content type
	if contentType != "post" && contentType != "page" {
		return fmt.Errorf("invalid content type: %s (must be 'post' or 'page')", contentType)
	}

	// Load config
	cfg, err := config.Load("config.toml")
	if err != nil {
		return fmt.Errorf("failed to load config (run 'ssg init' first): %w", err)
	}

	// Determine output directory
	var outputDir string
	if contentType == "post" {
		outputDir = filepath.Join(cfg.Build.ContentDir, "posts")
	} else {
		outputDir = filepath.Join(cfg.Build.ContentDir, "pages")
	}

	// Ensure directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create filename (sanitize name)
	filename := strings.ReplaceAll(strings.ToLower(name), " ", "-")
	if !strings.HasSuffix(filename, ".md") {
		filename += ".md"
	}
	filePath := filepath.Join(outputDir, filename)

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("file already exists: %s", filePath)
	}

	// Create content with frontmatter
	content := generateContent(contentType, name)

	// Write file
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	fmt.Printf("✓ Created new %s: %s\n", contentType, filePath)
	return nil
}

func generateContent(contentType, name string) string {
	now := time.Now().Format("2006-01-02")

	// Base frontmatter
	frontmatter := fmt.Sprintf(`+++
title = "%s"
date = "%s"
draft = false
`, name, now)

	// Add type-specific frontmatter
	if contentType == "post" {
		frontmatter += `author = ""
description = ""
`
	}

	frontmatter += "+++\n\n"

	// Add template content
	if contentType == "post" {
		frontmatter += fmt.Sprintf(`# %s

Write your post content here...

## Section 1

Add your content using Markdown syntax.

`, name)
	} else {
		frontmatter += fmt.Sprintf(`# %s

Write your page content here...

`, name)
	}

	return frontmatter
}
