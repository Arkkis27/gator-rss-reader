package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

// Config holds the application configuration
type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	UserAgent       string `json:"user_agent"`
}

// LoadConfig reads and returns the configuration from disk
func LoadConfig() (*Config, error) {
	cfg, err := read()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Expand environment variables in the DB URL
	expandedURL := os.ExpandEnv(cfg.DBUrl)

	// Parse and re-encode the URL to ensure proper escaping
	u, err := url.Parse(expandedURL)
	if err != nil {
		return nil, fmt.Errorf("invalid database URL: %w", err)
	}
	cfg.DBUrl = u.String()

	return &cfg, nil
}

// SetUser updates the current user in the config and saves it
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}

// private helper functions below

func getConfigFilePath() (string, error) {
	// Get the file path to the config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := homeDir + "/.gatorconfig.json"
	return fullPath, nil
}

func write(cfg Config) error {
	// Get the file path to the config
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// Convert the Config struct to a JSON byte slice
	jsonData, err := json.MarshalIndent(cfg, "", "  ") // Use MarshalIndent for pretty formatting
	if err != nil {
		return err
	}
	// Write the JSON to the file
	err = os.WriteFile(path, jsonData, 0644) // 0644 means read-write for owner, read-only for group/others
	if err != nil {
		return err
	}
	return nil
}

func read() (Config, error) {
	var cfg Config

	// Get the file path to the config
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err // Return on error
	}

	// Read the file contents
	dat, err := os.ReadFile(path)
	if err != nil {
		// Create default config if file doesn't exist
		err = createDefaultConfig()
		if err != nil {
			return cfg, fmt.Errorf("failed to create default config: %w", err)
		}
		// Update user what happened
		fmt.Println("No config found, default config saved at ~/.gatorconfig.json")
		fmt.Println("Update the DBUrl in the config to run the program")

		// Try to read the newly created config
		dat, err = os.ReadFile(path)
		if err != nil {
			return cfg, fmt.Errorf("failed to read new config: %w", err)
		}
	}

	// Decode JSON into the Config struct
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return cfg, err // Return on error
	}

	return cfg, nil // Return the populated cfg struct
}

func createDefaultConfig() error {
	defaultCfg := Config{
		DBUrl:           "postgres://gator_user:${GATOR_DB_PASSWORD}@<your_db_server_ip>:5432/gator_db?sslmode=disable",
		CurrentUserName: "",
		UserAgent:       "Gator-RSS-Reader/1.0",
	}
	return write(defaultCfg)
}
