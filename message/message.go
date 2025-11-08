package message

type Message interface {
	ID() uint32
	SetID(id uint32)
	IsServerEvent() bool
}

type message struct {
	id uint32
}

func (m message) ID() uint32 {
	return m.id
}

func (m *message) SetID(id uint32) {
	m.id = id
}

func (m message) IsServerEvent() bool {
	return m.id == 0
}