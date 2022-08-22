package invoker

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// 响应报文
type data struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

// 封装第三方调用服务
func (i *Invoker) PostForm(path string, params url.Values, result interface{}) error {
	params.Add("api_key", i.config.AppKey[path])
	resp, err := http.PostForm("https://api.muxiaoguo.cn/api"+path, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d := new(data)
	if err := json.NewDecoder(resp.Body).Decode(d); err != nil {
		return err
	}
	return json.Unmarshal(d.Data, result)
}
