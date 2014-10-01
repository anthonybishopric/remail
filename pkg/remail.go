package remail

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/go-smtpd/smtpd"
)

type OnEmail func(r ReceivedEmail, jsonContent []byte) error

func Serve(listen string, onEmail OnEmail) error {

	server := smtpd.Server{
		Addr:      listen,
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
