package main

import (
  //"bytes"
  //"context"
  //"encoding/gob"
  "fmt"
  /*
     "github.com/Shopify/sarama"
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
  //  then it calls email function???
  // Sender data.
  from := os.Getenv("SENDER_EMAIL")
  password := os.Getenv("SENDER_PASSWORD")
  key := os.Getenv("SENDGRID_API_KEY")

  conf := &EmailConfig{
    ApiKey:   key,
    From:     from,
    Password: password,
  }

  err := SendApiEmail(&EmailAccountEvent{BuyerEmail: "dagodking111@gmail.com"}, conf)

}
