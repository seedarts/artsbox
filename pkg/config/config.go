package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Settings 应用配置结构
type Settings struct {
	APIKey  string `json:"api_key"`
	BaseURL string `json:"base_url"`
	Model   string `json:"model"`
	SaveDir string `json:"save_dir"`
}

// DefaultSettings 返回默认配置
func DefaultSettings() *Settings {
	homeDir, _ := os.UserHomeDir()
	return &Settings{
		APIKey:  "",
		BaseURL: "https://api.cozex.cn/v1",
		Model:   "doubao-seedream-4-5-251128",
		SaveDir: filepath.Join(homeDir, "Pictures", "ArtsBox"),
	}
}

// GetConfigPath 返回配置文件路径
func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".artsbox")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}

// Load 从文件加载配置
func Load() (*Settings, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return DefaultSettings(), err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultSettings(), nil
		}
		return DefaultSettings(), err
	}

	var settings Settings
	if err := json.Unmarshal(data, &settings); err != nil {
		return DefaultSettings(), err
	}

	// 确保保存目录存在
	if settings.SaveDir == "" {
		settings.SaveDir = DefaultSettings().SaveDir
	}
	os.MkdirAll(settings.SaveDir, 0755)

	return &settings, nil
}

// Save 保存配置到文件
func Save(settings *Settings) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// 确保保存目录存在
	if settings.SaveDir != "" {
		os.MkdirAll(settings.SaveDir, 0755)
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
