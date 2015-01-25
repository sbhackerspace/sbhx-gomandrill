// Steve Phillips / elimisteve
// 2014.03.29

package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API URLs
const (
	baseURL = "https://mandrillapp.com/api/1.0"

	messagesSendURL = baseURL + "/messages/send.json"
)

type MandrillMessageSender struct {
	Key string `json:"key"`
}

func NewSender(key string) (*MandrillMessageSender, error) {
	if key == "" {
		return nil, fmt.Errorf("key cannot be blank\n")
	}
	return &MandrillMessageSender{Key: key}, nil
}

func (sender *MandrillMessageSender) SendSimple(from string, to []string, subject, body string) error {
	m := &Message{
		Html:      body,
		Text:      body,
		Subject:   subject,
		FromEmail: from,
		To:        emailsToRecipients(to),
	}
	return sender.Send(m, false)
}

func (sender *MandrillMessageSender) Send(msg *Message, async bool) error {
	m := &MandrillMessage{
		Key:     sender.Key,
		Message: msg,
		Async:   async,
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("Error marshaling *MandrillMessage: %v\n", err)
	}

	reader := bytes.NewReader(jsonData)
	resp, err := http.Post(messagesSendURL, "application/json", reader)
	if err != nil {
		return fmt.Errorf("Error sending *MandrillMessage: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
	}

	return ValidateSendResponses(body)
}
