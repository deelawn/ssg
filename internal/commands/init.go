package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deelawn/ssg/internal/config"
	"github.com/deelawn/ssg/internal/embedtheme"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialize a new static site",
	Long:  `Create a new static site with default structure and configuration.`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	// Determine target directory
	targetDir := "."
	if len(args) > 0 {
		targetDir = args[0]
	}

	// Create target directory if it doesn't exist
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create directory structure
	dirs := []string{
		"content/posts",
		"content/pages",
		"themes/default/templates",
		"themes/default/static",
		"public",
	}

	for _, dir := range dirs {
		path := filepath.Join(targetDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create default config
	cfg := config.Default()
	configPath := filepath.Join(targetDir, "config.toml")
	if err := config.Save(configPath, cfg); err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}

	// Create sample content
	samplePost := `+++
title = "Welcome to Your New Site"
date = "2024-01-01"
draft = false
+++

# Welcome!

This is your first blog post. Edit this file to customize it.

## Getting Started

- Run ` + "`ssg add post my-post`" + ` to create a new post
- Run ` + "`ssg add page about`" + ` to create a new page
- Run ` + "`ssg build`" + ` to generate your site
`

	postPath := filepath.Join(targetDir, "content/posts/welcome.md")
	if err := os.WriteFile(postPath, []byte(samplePost), 0644); err != nil {
		return fmt.Errorf("failed to create sample post: %w", err)
	}

	// Create .gitignore
	gitignore := `public/
.DS_Store
`
	gitignorePath := filepath.Join(targetDir, ".gitignore")
	if err := os.WriteFile(gitignorePath, []byte(gitignore), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	// Copy default theme files
	if err := createDefaultTheme(targetDir); err != nil {
		return fmt.Errorf("failed to create default theme: %w", err)
	}

	fmt.Printf("✓ Initialized new static site in %s\n", targetDir)
	fmt.Println("✓ Created directory structure")
	fmt.Println("✓ Created config.toml")
	fmt.Println("✓ Created sample content")
	fmt.Println("✓ Created default theme")
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Edit config.toml to customize your site")
	fmt.Println("  2. Run 'ssg build' to generate your site")
	fmt.Println("  3. Run 'ssg add post <name>' to create new posts")

	return nil
}

func createDefaultTheme(targetDir string) error {
	themeDir := filepath.Join(targetDir, "themes/default")
	return embedtheme.WriteThemeFiles(themeDir)
}
