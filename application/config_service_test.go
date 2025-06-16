package application

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestConfigService_DefaultConfig(t *testing.T) {
	config := DefaultConfig()
	
	if config.WorkDuration != 25*time.Minute {
		t.Errorf("Expected work duration 25m, got %v", config.WorkDuration)
	}
	
	if config.BreakDuration != 5*time.Minute {
		t.Errorf("Expected break duration 5m, got %v", config.BreakDuration)
	}
	
	if config.WarningInterval != 5*time.Minute {
		t.Errorf("Expected warning interval 5m, got %v", config.WarningInterval)
	}
	
	if !config.SoundEnabled {
		t.Error("Sound should be enabled by default")
	}
	
	if config.Volume != 0.7 {
		t.Errorf("Expected default volume 0.7, got %v", config.Volume)
	}
}

func TestConfigService_NewConfigService(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "karedoro_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Override home directory for test
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", originalHome)
	
	service := NewConfigService()
	
	config := service.GetConfig()
	if config == nil {
		t.Error("Config should not be nil")
	}
	
	// Should have default values
	if config.WorkDuration != 25*time.Minute {
		t.Errorf("Expected work duration 25m, got %v", config.WorkDuration)
	}
}

func TestConfigService_SaveAndLoad(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "karedoro_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Override home directory for test
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", originalHome)
	
	service := NewConfigService()
	
	// Modify config
	newConfig := &Config{
		WorkDuration:    30 * time.Minute,
		BreakDuration:   10 * time.Minute,
		WarningInterval: 3 * time.Minute,
		SoundEnabled:    false,
		Volume:          0.5,
	}
	
	err = service.UpdateConfig(newConfig)
	if err != nil {
		t.Errorf("UpdateConfig should not return error, got %v", err)
	}
	
	// Create new service to test loading
	service2 := NewConfigService()
	loadedConfig := service2.GetConfig()
	
	if loadedConfig.WorkDuration != 30*time.Minute {
		t.Errorf("Expected loaded work duration 30m, got %v", loadedConfig.WorkDuration)
	}
	
	if loadedConfig.BreakDuration != 10*time.Minute {
		t.Errorf("Expected loaded break duration 10m, got %v", loadedConfig.BreakDuration)
	}
	
	if loadedConfig.WarningInterval != 3*time.Minute {
		t.Errorf("Expected loaded warning interval 3m, got %v", loadedConfig.WarningInterval)
	}
	
	if loadedConfig.SoundEnabled {
		t.Error("Sound should be disabled")
	}
	
	if loadedConfig.Volume != 0.5 {
		t.Errorf("Expected loaded volume 0.5, got %v", loadedConfig.Volume)
	}
}

func TestConfigService_ConfigFileCreation(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "karedoro_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Override home directory for test
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", originalHome)
	
	// Create service (should create config file)
	_ = NewConfigService()
	
	// Check if config file was created
	expectedPath := filepath.Join(tempDir, ".karedoro", "config.json")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Error("Config file should have been created")
	}
	
	// Check if config directory was created
	expectedDir := filepath.Join(tempDir, ".karedoro")
	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		t.Error("Config directory should have been created")
	}
}

func TestConfigService_SaveError(t *testing.T) {
	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "karedoro_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Create a file where directory should be (to cause error)
	configPath := filepath.Join(tempDir, ".karedoro")
	err = os.WriteFile(configPath, []byte("not a directory"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	
	// Override home directory for test
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", originalHome)
	
	service := NewConfigService()
	
	// Save should fail because .karedoro is a file, not a directory
	newConfig := DefaultConfig()
	newConfig.Volume = 0.9
	
	err = service.UpdateConfig(newConfig)
	if err == nil {
		t.Error("UpdateConfig should return error when unable to create directory")
	}
}