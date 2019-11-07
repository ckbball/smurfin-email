// only events needed to handle is subscribe to EmailAccountEvent which triggers emailing process
package main

import (
  "bytes"
  "context"
  "encoding/gob"
  "github.com/ThreeDotsLabs/watermill"
  "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
  "github.com/ThreeDotsLabs/watermill/message"
  catalogProto "github.com/ckbball/smurfin-catalog/proto/catalog"
  pb "github.com/ckbball/smurfin-checkout/proto/checkout"
  "time"
)

type EmailAccountEvent struct {
  BuyerId              string
  AccountLogin         string
  AccountPassword      string
  AccountEmail         string
  AccountEmailPassword string
}

func InitSubscriber(config kafka.SubscriberConfig) *kafka.Subscriber {
  subscriber, err := kafka.NewSubscriber(
    kafka.SubscriberConfig{
      Brokers:               []string{"kafka:9092"},
      Unmarshaler:           kafka.DefaultMarshaler{},
      OverwriteSaramaConfig: config,
      ConsumerGroup:         "test_consumer_group",
    },
    watermill.NewStdLogger(false, false),
  )
  if err != nil {
    panic(err)
  }
  return subscriber
}
