package remail

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bradfitz/go-smtpd/smtpd"
)

type OnEmail func(r ReceivedEmail, jsonContent []byte) error

func Serve(onEmail OnEmail) error {

	server := smtpd.Server{
		Addr:      PortNumber(),
		OnNewMail: ListenWithOnEmail(onEmail),
	}
	return server.ListenAndServe()
}

type ReceivedEmail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Addr    string   `json:"addr"`
	Content string   `json:"content"`
}

func PortNumber() string {
	port := os.Getenv("SMTP_PORT")
	if port == "" {
		port = ":25"
	}
	return port
}

type RemailEnvelope struct {
	onEmail OnEmail
	rec     ReceivedEmail
}

func (r *RemailEnvelope) AddRecipient(rcpt smtpd.MailAddress) error {
	r.rec.To = append(r.rec.To, rcpt.Email())
	return nil
}

func (r *RemailEnvelope) BeginData() error {
	return nil
}

func (r *RemailEnvelope) Write(v []byte) error {
	r.rec.Content += string(v)
	return nil
}

func (r *RemailEnvelope) Close() error {
	jsonContent, err := json.Marshal(r.rec)
	log.Println(string(jsonContent))
	if err != nil {
		return err
	}
	return r.onEmail(r.rec, jsonContent)
}

func ListenWithOnEmail(onEmail OnEmail) func(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	return func(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
		rec := ReceivedEmail{
			From: from.Email(),
			Addr: c.Addr().String(),
		}
		return &RemailEnvelope{
			onEmail: onEmail,
			rec:     rec,
		}, nil
	}
}

func ListenWithPrint(r ReceivedEmail, jsonContent []byte) error {
	fmt.Println(string(jsonContent))
	return nil
}
