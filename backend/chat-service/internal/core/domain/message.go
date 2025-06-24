package domain

import "time"

type Message struct {
	ID        int
	RoomID    int
	Content   string
	SenderID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
