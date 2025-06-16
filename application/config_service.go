package application

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	WorkDuration    time.Duration `json:"work_duration"`
	BreakDuration   time.Duration `json:"break_duration"`
	WarningInterval time.Duration `json:"warning_interval"`
	SoundEnabled    bool          `json:"sound_enabled"`
	Volume          float64       `json:"volume"`
}

func DefaultConfig() *Config {
	return &Config{
		WorkDuration:    25 * time.Minute,
		BreakDuration:   5 * time.Minute,
		WarningInterval: 5 * time.Minute,
		SoundEnabled:    true,
		Volume:          0.7,
	}
}

type ConfigService struct {
	config     *Config
	configPath string
}

func NewConfigService() *ConfigService {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".karedoro", "config.json")
	
	service := &ConfigService{
		config:     DefaultConfig(),
		configPath: configPath,
	}
	
	service.Load()
	return service
}

func (c *ConfigService) Load() error {
	if _, err := os.Stat(c.configPath); os.IsNotExist(err) {
		return c.Save()
	}
	
	data, err := os.ReadFile(c.configPath)
	if err != nil {
		return err
	}
	
	return json.Unmarshal(data, c.config)
}

func (c *ConfigService) Save() error {
	dir := filepath.Dir(c.configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	data, err := json.MarshalIndent(c.config, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(c.configPath, data, 0644)
}

func (c *ConfigService) GetConfig() *Config {
	return c.config
}

func (c *ConfigService) UpdateConfig(config *Config) error {
	c.config = config
	return c.Save()
}