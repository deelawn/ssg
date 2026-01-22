package commands

import (
	"fmt"
	"time"

	"github.com/deelawn/ssg/internal/config"
	"github.com/deelawn/ssg/internal/generator"
	"github.com/deelawn/ssg/internal/theme"
	"github.com/spf13/cobra"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the static site",
	Long:  `Convert markdown files to HTML and generate the complete static site.`,
	RunE:  runBuild,
}

func runBuild(cmd *cobra.Command, args []string) error {
	startTime := time.Now()

	// Load configuration
	cfg, err := config.Load("config.toml")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	fmt.Println("Building site...")
	fmt.Printf("Theme: %s\n", cfg.Theme.Name)

	// Load theme
	thm, err := theme.Load(cfg.Theme.Dir, cfg.Theme.Name)
	if err != nil {
		return fmt.Errorf("failed to load theme: %w", err)
	}

	// Create generator
	gen := generator.New(cfg, thm)

	// Build site
	if err := gen.Build(); err != nil {
		return err
	}

	duration := time.Since(startTime)
	fmt.Printf("\n✓ Site built successfully in %v\n", duration.Round(time.Millisecond))
	fmt.Printf("✓ Output directory: %s\n", cfg.Build.OutputDir)

	return nil
}
