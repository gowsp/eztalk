package wss

import "fmt"

type opcode int

const (
	Dispatch      opcode = 0
	Heartbeat     opcode = 1
	Identify      opcode = 2
	Hello         opcode = 10
	Heartbeat_ACK opcode = 11
)

func (o opcode) Event(val string) string {
	return fmt.Sprintf("%d:%s", o, val)
}
