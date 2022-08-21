package wss

import (
	"context"
	"encoding/json"
	"fmt"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type event struct {
	Type   string          `json:"t,omitempty"`
	OpCode int             `json:"op,omitempty"`
	Seq    uint64          `json:"s,omitempty"`
	Data   json.RawMessage `json:"d,omitempty"`
}

func (e *event) key() string {
	return fmt.Sprintf("%d:%s", e.OpCode, e.Type)
}

type output struct {
	OpCode opcode      `json:"op,omitempty"`
	Data   interface{} `json:"d,omitempty"`
}

func write(conn *websocket.Conn, code opcode, data interface{}) error {
	msg := output{
		OpCode: code,
		Data:   data,
	}
	d, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(string(d))
	return wsjson.Write(context.Background(), conn, msg)
}
