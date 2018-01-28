package shared

import "time"

type ChatMessage struct {
	Name    string
	Time    time.Time
	Message string
}
