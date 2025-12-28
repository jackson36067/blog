package models

import "time"

type Conversation struct {
	Model
	Type            string    `json:"type"`
	LastMessage     string    `json:"last_message"`
	LastMessageTime time.Time `json:"last_message_time"`
}
