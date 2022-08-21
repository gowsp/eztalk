package invoker

// 创建配置
func NewConfig(appId, token string) Config {
	return Config{
		AppId: appId,
		Token: token,
	}
}

// 配置信息
type Config struct {
	Env    string `json:"env,omitempty"`
	AppId  string `json:"bot_app_id,omitempty"`
	Secret string `json:"bot_secret,omitempty"`
	Token  string `json:"bot_token,omitempty"`
}

// 认证
func (c *Config) BotToken() string {
	return c.AppId + "." + c.Token
}
