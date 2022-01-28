package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"os"
	"os/signal"
)

func ConnectAndListening(c chan []byte) {
	var URL string
	sc, _:= stan.Connect("test-cluster", "myID")
	nc, _ := nats.Connect(URL)
	defer nc.Close()
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", m.Data)
		c <- m.Data
	}, stan.DeliverAllAvailable())

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		for range signalChan {
			sub.Close()
			sc.Close()
			close(c)
			cleanupDone <- true
		}
	}()
	<-cleanupDone
	sub.Close()
}
