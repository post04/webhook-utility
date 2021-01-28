package discordhook

import (
	"bytes"
	"fmt"
	"net/http"
)

// WebhookPayload - the payload that is given to the send function
type WebhookPayload struct {
	Content   string
	Username  string
	AvatarURL string
}

// Send - send a message via webhook
func Send(webhook string, payload WebhookPayload) error {
	if payload.Content == "" {
		return fmt.Errorf("\"Content\" not found")
	}
	jsonstr := fmt.Sprintf("{\"content\": \"%s\"", payload.Content)
	if payload.AvatarURL != "" {
		jsonstr += fmt.Sprintf(", \"avatar_url\": \"%s\"", payload.AvatarURL)
	}
	if payload.Username != "" {
		jsonstr += fmt.Sprintf(", \"username\": \"%s\"", payload.Username)
	}
	jsonstr += "}"
	tosend := []byte(jsonstr)
	client := &http.Client{}
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(tosend))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
