package config

import (
	"gopkg.in/ini.v1"
)

// Config 存储应用配置
type Config struct {
	Deepseek DeepseekConfig
}

// DeepseekConfig 存储 deepseek 相关配置
type DeepseekConfig struct {
	APIKey string
	URL    string
	Model  string
}

// AppConfig 全局配置实例
var AppConfig Config

// LoadConfig 加载配置文件
func LoadConfig() error {
	cfg, err := ini.Load("./init/project.ini")
	if err != nil {
		return err
	}

	// 加载 deepseek 配置
	AppConfig.Deepseek = DeepseekConfig{
		APIKey: cfg.Section("deepseek").Key("api_key").String(),
		URL:    cfg.Section("deepseek").Key("url").String(),
		Model:  cfg.Section("deepseek").Key("model").String(),
	}

	return nil
}
