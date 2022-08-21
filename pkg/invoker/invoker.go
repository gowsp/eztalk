package invoker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var Err_Response = errors.New("error status")

// 创建调用器
func New(config *Config) *Invoker {
	token := fmt.Sprintf("Bot %s.%s", config.AppId, config.Token)
	return &Invoker{
		config: config,
		token:  token,
	}
}

// 调用器，业务无关，封装对接口调用
type Invoker struct {
	config *Config
	token  string
}

// 实际调用
func (i *Invoker) do(req *http.Request, result interface{}) error {
	req.Header.Set("Authorization", i.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Err_Response
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

// 机器人Token
func (i *Invoker) BotToken() string {
	return i.token
}

// Get请求
func (i *Invoker) Get(path string, result interface{}) error {
	req, err := http.NewRequest(http.MethodGet, i.config.Env+path, nil)
	if err != nil {
		return err
	}
	return i.do(req, result)
}

// Post请求
func (i *Invoker) Post(path string, body, result interface{}) error {
	val, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, i.config.Env+path, bytes.NewBuffer(val))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	return i.do(req, result)
}
