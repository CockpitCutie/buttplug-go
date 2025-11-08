package message

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Message interface {
	ID() uint32
	SetID(id uint32)
	IsServerEvent() bool
}

type message struct {
	Id uint32
}

func (m message) ID() uint32 {
	return m.Id
}

func (m *message) SetID(id uint32) {
	m.Id = id
}

func (m message) IsServerEvent() bool {
	return m.Id == 0
}

func Serialize(m Message) (string, error) {
	messageKind := reflect.TypeOf(m).Elem().Name()
	serialized, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`[{"%s":%s}]`, messageKind, string(serialized)), nil
}
