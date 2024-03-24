package models

type ChatMessage struct {
	Type     int    `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
