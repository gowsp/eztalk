package invoker

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
