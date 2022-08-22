package cmd

import (
	"log"
	"net/url"
	"sync"

	"github.com/eztalk/pkg/invoker"
)

var cmds = map[string]TopicType{
	"/歇后语":   Fable,
	"/猜字谜":   Riddles,
	"/脑筋急转弯": BrainTwists,
}

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

// 趣味问答
type topicCmd struct {
	session sync.Map
	invoker *invoker.Invoker
}

// 消息ID
func (c *topicCmd) Id() []string {
	return []string{"/歇后语", "/猜字谜", "/脑筋急转弯"}
}

// 交互型，需记录会话
func (c *topicCmd) Interactive() bool {
	return true
}

// 题目
type topic struct {
	Type   TopicType
	answer *answer
	UUID   string `json:"uuid"`
	Riddle string `json:"riddle"`
}

func (c *topicCmd) Handle(input Input) (string, error) {
	topicType := cmds[input.UserInput()]
	params := make(url.Values)
	params.Add("type", string(topicType))
	params.Add("output", "riddle")
	question := new(topic)
	if err := c.invoker.PostForm("/amusingQA", params, question); err != nil {
		return "", err
	}
	question.Type = topicType
	c.session.Store(input.UserId(), question)
	return question.Riddle, nil
}

func (s *topicCmd) Reply(input Input) string {
	userId := input.UserId()
	data, ok := s.session.Load(userId)
	if !ok {
		return "服务开小差了"
	}
	answer, err := s.getAnswer(data.(*topic))
	if err != nil {
		log.Println(err)
		return "服务开小差了"
	}
	request := input.UserInput()
	if request == "答案" {
		s.session.Delete(userId)
		return "答案：" + answer.Answer
	}
	if answer.isCorrect(request) {
		s.session.Delete(userId)
		return "厉害啊，恭喜客官答对了"
	}
	return "客官回答错误了，再试试"
}

// 回答
type answer struct {
	Answer string `json:"answer"`
}

// 检查是否正确
func (a *answer) isCorrect(input string) bool {
	return a.Answer == input
}

// 获取答案
func (s *topicCmd) getAnswer(t *topic) (*answer, error) {
	if t.answer != nil {
		return t.answer, nil
	}
	params := make(url.Values)
	params.Add("type", string(t.Type))
	params.Add("output", "answer")
	params.Add("uuid", t.UUID)
	answer := new(answer)
	if err := s.invoker.PostForm("/amusingQA", params, answer); err != nil {
		return nil, err
	}
	t.answer = answer
	return answer, nil
}
