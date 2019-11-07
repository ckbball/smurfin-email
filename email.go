package main

import (
  "fmt"
  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
  "log"
)

func SendSMTPEmail(email string, password string, accountInfo EmailAccountEvent) error {

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

func SendApiEmail(account EmailAccountEvent, config EmailConfig) error {
  from := mail.NewEmail("Example User", config.From)
  subject := "Sending with SendGrid is Fun"
  to := mail.NewEmail("Example User", account.BuyerEmail)
  plainTextContent := "and easy to do anywhere, even with Go"
  htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
  message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
  client := sendgrid.NewSendClient(config.ApiKey)
  response, err := client.Send(message)
  if err != nil {
    log.Println(err)
    return err
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
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

type EmailConfig struct {
  ApiKey   string
  From     string
  Password string
}
