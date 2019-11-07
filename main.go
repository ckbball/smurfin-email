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
  // Receiver email address.
  to := []string{
    "blah1@gmail.com",
    "blah2@gmail.com",
  }
  // smtp server configuration.
  smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
  // Message.
  message := []byte("This is a really unimaginative message, I know.")
  // Authentication.
  auth := smtp.PlainAuth("", from, password, smtpServer.host)
  // Sending email.
  err := smtp.SendMail(smtpServer.serverName(), auth, from, to, message)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("Email Sent!")
}

// smtpServer data to smtp server
type smtpServer struct {
  host string
  port string
}

// serverName URI to smtp server
func (s *smtpServer) serverName() string {
  return s.host + ":" + s.port
}
