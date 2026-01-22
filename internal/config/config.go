package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config represents the site configuration
type Config struct {
	Site  SiteConfig  `toml:"site"`
	Build BuildConfig `toml:"build"`
	Theme ThemeConfig `toml:"theme"`
}

// SiteConfig contains site-level settings
type SiteConfig struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	Domain      string `toml:"domain"`
	Author      string `toml:"author"`
}

// BuildConfig contains build settings
type BuildConfig struct {
	ContentDir string `toml:"content_dir"`
	OutputDir  string `toml:"output_dir"`
}

// ThemeConfig contains theme settings
type ThemeConfig struct {
	Name string `toml:"name"`
	Dir  string `toml:"dir"`
}

// Default returns a default configuration
func Default() *Config {
	return &Config{
		Site: SiteConfig{
			Title:       "My Static Site",
			Description: "A site generated with SSG",
			Domain:      "example.com",
			Author:      "",
		},
		Build: BuildConfig{
			ContentDir: "content",
			OutputDir:  "public",
		},
		Theme: ThemeConfig{
			Name: "default",
			Dir:  "themes",
		},
	}
}

// Load reads configuration from a TOML file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Save writes configuration to a TOML file
func Save(path string, cfg *Config) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
