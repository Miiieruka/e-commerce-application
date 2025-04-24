package repository

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
