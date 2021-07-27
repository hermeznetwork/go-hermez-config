package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func loadDefault(defaultValues string, cfg interface{}) error {
	if _, err := toml.Decode(defaultValues, cfg); err != nil {
		return err
	}
	return nil
}
func loadEnv(cfg interface{}) error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}
func loadFile(path string, cfg interface{}) error {
	bs, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return err
	}
	cfgToml := string(bs)
	if _, err := toml.Decode(cfgToml, cfg); err != nil {
		return err
	}
	return nil
}

//LoadConfig is the function that loads the configuration
func LoadConfig(filePath string, defaultValues string, cfg interface{}) error {
	//Get default configuration
	if err := loadDefault(defaultValues, cfg); err != nil {
		return fmt.Errorf("error loading default configuration: %w", err)
	}
	// Get file configuration
	if err := loadFile(filePath, cfg); err != nil {
		return fmt.Errorf("error loading configuration file: %w", err)
	}
	// Overwrite file configuration with the env configuration
	if err := loadEnv(cfg); err != nil {
		return fmt.Errorf("error getting configuration from env: %w", err)
	}
	return nil
}