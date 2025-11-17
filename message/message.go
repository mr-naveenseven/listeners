package message

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type Message struct {
	data      string
	stopper   chan struct{}
	wg        *sync.WaitGroup
	msgLocker sync.Mutex
}

func NewMessage(wg *sync.WaitGroup) *Message {
	return &Message{
		data:    "initial message",
		stopper: make(chan struct{}),
		wg:      wg,
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

func (m *Message) messageUpdater() {
	m.wg.Add(1)
	defer m.wg.Done()
	for i := 0; ; i++ {
		select {
		case <-m.stopper:
			log.Println("Message updater stopped")
			return
		default:
			m.writeMessage("message number " + strconv.Itoa(i))
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func (m *Message) StartUpdating() {
	go m.messageUpdater()
}

func (m *Message) StopUpdating() {
	close(m.stopper)
}

// TODO: Add more methods to manipulate message if needed
