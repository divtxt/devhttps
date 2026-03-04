package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Entry struct {
	Domain string `json:"domain"`
	Port   int    `json:"port"`
}

type Config struct {
	Entries []Entry `json:"entries"`
}

func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".devhttps", "config.json"), nil
}

func Load() (*Config, error) {
	path, err := Path()
	if err != nil {
		return &Config{}, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Config{}, nil
	}
	if err != nil {
		return &Config{}, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return &Config{}, err
	}
	return &cfg, nil
}

func Save(c *Config) error {
	path, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func (c *Config) Add(domain string, port int) {
	for i, e := range c.Entries {
		if e.Domain == domain {
			c.Entries[i].Port = port
			return
		}
	}
	c.Entries = append(c.Entries, Entry{Domain: domain, Port: port})
}
