package main

import (
  //"bytes"
  //"context"
  //"encoding/gob"
  "fmt"
  "github.com/Shopify/sarama"
  /*
     "github.com/ThreeDotsLabs/watermill"
     "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
     "github.com/ThreeDotsLabs/watermill/message"
     catalogProto "github.com/ckbball/smurfin-catalog/proto/catalog"
     pb "github.com/ckbball/smurfin-checkout/proto/checkout"
  */
  "github.com/subosito/gotenv"
  "net/smtp"
  "os"
  //"time"
)

func init() {
  gotenv.Load()
}

func main() {
  // kafka subscriber
  // user-service client
  // run the subscriber
  // then it calls email function???
  // Sender data.
  from := os.Getenv("SENDER_EMAIL")
  password := os.Getenv("SENDER_PASSWORD")
  key := os.Getenv("SENDGRID_API_KEY")

  conf := &EmailConfig{
    ApiKey:   key,
    From:     from,
    Password: password,
  }
  // Make subscriber config here
  saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
  saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

  // Make subscriber pointer here
  subscriber := InitSubscriber(saramaSubscriberConfig)

  // Make user service client
  service := micro.NewService(micro.Name("smurfin.email"))
  service.Init()

  userClient := userProto.NewUserServiceClient("smurfin.user.client", service.Client())

  go process(subscriber, conf, userClient)

}

func process(sub *kafka.Subscriber, conf *EmailConfig, client *UserServiceClient) {
  // listen for messages on the emailtopic
  messages, err := s.subscriber.Subscribe(context.Background(), "email.topic")
  if err != nil {
    log.Printf(err)
  }
  // when a message is received:
  for msg := range messages {
    // decode msg payload back into struct
    var network bytes.Buffer
    var ea EmailAccountEvent
    network.Write(msg.payload)
    dec := gob.NewDecoder(&network)
    err = dec.Decode(&ea)
    if err != nil {
      log.Fatal("decode error: ", err)
    }
    log.Printf("received message: %s, payload buyer_id: %s", msg.UUID, ea.BuyerId)
    log.Printf("Checking if correct payload received. buyer id: %s || account: %s", buyer_id, accountId)

    // Make API call here to user
    // returns a user.User object
    cr, err := client.FindUser(context.Context, &userProto.Request{
      UserId: ea.BuyerId,
    })
    log.Printf("Found user with id: %s \n", ea.BuyerId)
    if err != nil {
      log.Printf("Error in user api call: ", err)
    }

    // Send email to user from message
    err = SendApiEmail(ea, conf)
    if err != nil {
      log.Printf("Error in sending email: ", err)
    }
  }
  // ?? Maybe confirm email address in the future

}
