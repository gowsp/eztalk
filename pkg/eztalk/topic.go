package eztalk

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// 题目类型
type TopicType string

const (
	//歇后语
	Fable TopicType = "fable"
	// 字谜
	Riddles TopicType = "riddles"
	//脑经急转弯
	BrainTwists TopicType = "brainTwists"
)

// 响应报文
type Data struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

// 题目
type Topic struct {
	Type   TopicType
	answer *Answer
	UUID   string `json:"uuid"`
	Riddle string `json:"riddle"`
}

// 回答
type Answer struct {
	Answer string `json:"answer"`
}

// 检查答案是否正确
func (a *Answer) isCorrect(answer string) bool {
	return a.Answer == answer
}

// 创建题目
func NewTopic(t TopicType) (*Topic, error) {
	params := make(url.Values)
	params.Add("api_key", "0cfb886d23ab1e83")
	params.Add("type", string(t))
	params.Add("output", "riddle")
	resp, err := http.PostForm("https://api.muxiaoguo.cn/api/amusingQA", params)
	question := new(Topic)
	err = unmarshal(resp, err, question)
	if err == nil {
		question.Type = t
	}
	return question, err
}

// 获取答案
func (t *Topic) GetAnswer() (*Answer, error) {
	if t.answer != nil {
		return t.answer, nil
	}
	params := make(url.Values)
	params.Add("api_key", "0cfb886d23ab1e83")
	params.Add("type", string(t.Type))
	params.Add("output", "answer")
	params.Add("uuid", t.UUID)
	resp, err := http.PostForm("https://api.muxiaoguo.cn/api/amusingQA", params)
	answer := new(Answer)
	err = unmarshal(resp, err, answer)
	if err == nil {
		t.answer = answer
	}
	return answer, err
}

// 解析结果
func unmarshal(resp *http.Response, err error, result interface{}) error {
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d := new(Data)
	if err := json.NewDecoder(resp.Body).Decode(d); err != nil {
		return err
	}
	return json.Unmarshal(d.Data, result)
}
