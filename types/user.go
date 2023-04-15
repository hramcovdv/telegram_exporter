package types

import (
	"fmt"
)

type User struct {
	ID       int64
	ChatID   int64
	Messages int64
	Symbols  int64
}

func NewUser(id int64, chatid int64) *User {
	return &User{
		ID:       id,
		ChatID:   chatid,
		Symbols:  0,
		Messages: 0,
	}
}

func (u *User) IncMessages(msg string) {
	symbols := len(msg)

	if symbols > 0 {
		u.Symbols += int64(symbols)
		u.Messages++
	}
}

func (u *User) GetMyIdMsg() string {
	return fmt.Sprintf("Your user ID: %d\nCurrent chat ID: %d",
		u.ID, u.ChatID)
}

func (u *User) GetMyStatsMsg() string {
	return fmt.Sprintf("You sent %d messages (%d symbols)",
		u.Messages, u.Symbols)
}
