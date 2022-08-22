package event

import (
	"context"
	"encoding/json"
	"fmt"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// 事件数据
type event struct {
	Type   string          `json:"t,omitempty"`
	OpCode int             `json:"op,omitempty"`
	Seq    uint64          `json:"s,omitempty"`
	Data   json.RawMessage `json:"d,omitempty"`
}

// 事件key
func (e *event) key() string {
	return fmt.Sprintf("%d:%s", e.OpCode, e.Type)
}

// 回复 websocket 消息体
type reply struct {
	OpCode opcode      `json:"op,omitempty"`
	Data   interface{} `json:"d,omitempty"`
}

// 写消息
func write(conn *websocket.Conn, code opcode, data interface{}) error {
	msg := reply{
		OpCode: code,
		Data:   data,
	}
	return wsjson.Write(context.Background(), conn, msg)
}
