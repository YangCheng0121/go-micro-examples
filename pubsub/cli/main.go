package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/pborman/uuid"
	proto "go-micro-examples/pubsub/proto"
	"time"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTimer(time.Second)

	for _ = range t.C {
		// crate new event
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing %v", err)
		}
	}
}

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
	)

	// parse command line
	service.Init()

	// create publisher
	pub1 := micro.NewEvent("example.topic.pubsub.1", service.Client())
	pub2 := micro.NewEvent("example.topic.pubsub.2", service.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)

	// block forever
	select {}
}
