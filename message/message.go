package message

import "sync"

type Message struct {
	data      string
	msgLocker sync.Mutex
}

func NewMessage() *Message {
	return &Message{
		data: "initial message",
	}
}

func (m *Message) ReadMessage() string {
	m.msgLocker.Lock()
	defer m.msgLocker.Unlock()
	return m.data
}

func (m *Message) writeMessage(data string) {
	m.msgLocker.Lock()
	defer m.msgLocker.Unlock()
	m.data = data
}

// TODO: Add more methods to manipulate message if needed
