package commands

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/deelawn/ssg/internal/config"
	"github.com/deelawn/ssg/internal/generator"
	"github.com/deelawn/ssg/internal/theme"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Build and serve the static site",
	Long:  `Build the static site and start a local web server on port 5115.`,
	RunE:  runServe,
}

func runServe(cmd *cobra.Command, args []string) error {
	// First, build the site
	fmt.Println("Building site...")

	// Load configuration
	cfg, err := config.Load("config.toml")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

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

	fmt.Printf("\n✓ Site built successfully\n")
	fmt.Printf("✓ Output directory: %s\n\n", cfg.Build.OutputDir)

	// Check if output directory exists
	if _, err := os.Stat(cfg.Build.OutputDir); os.IsNotExist(err) {
		return fmt.Errorf("output directory does not exist: %s", cfg.Build.OutputDir)
	}

	// Start web server
	port := "5115"
	addr := fmt.Sprintf(":%s", port)

	// Create file server
	fs := http.FileServer(http.Dir(cfg.Build.OutputDir))
	http.Handle("/", fs)

	fmt.Printf("Starting web server...\n")
	fmt.Printf("Server running at http://localhost:%s\n", port)
	fmt.Printf("Press Ctrl+C to stop\n\n")

	// Start server
	server := &http.Server{
		Addr:         addr,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
