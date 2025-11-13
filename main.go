package main

import (
	"listeners/listener"
	"listeners/message"
	"sync"
	"time"
)

func main() {
	message := message.NewMessage()
	var wg sync.WaitGroup

	// Create and start listeners
	listener1 := listener.NewListener("Pro plan listener", time.Second*1, message, &wg)
	listener1.Start()
	listener2 := listener.NewListener("Base plan listener", time.Second*5, message, &wg)
	listener2.Start()

	time.Sleep(time.Second * 6)

	// Stop listeners
	listener1.Stop()
	listener2.Stop()

	wg.Wait()
}
