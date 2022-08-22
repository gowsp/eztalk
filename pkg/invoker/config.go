package invoker

import (
	"encoding/json"
	"os"
)

// 创建配置
func NewConfig(appId, token string) Config {
	return Config{
		AppId: appId,
		Token: token,
	}
}

// 从文件中读取配置
func ReadConfig(path string) (*Config, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}

// 配置信息
type Config struct {
	Env    string            `json:"bot_env,omitempty"`
	AppKey map[string]string `json:"app_key,omitempty"`
	AppId  string            `json:"bot_app_id,omitempty"`
	Secret string            `json:"bot_secret,omitempty"`
	Token  string            `json:"bot_token,omitempty"`
}

// 认证
func (c *Config) BotToken() string {
	return c.AppId + "." + c.Token
}
