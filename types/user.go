package types

type User struct {
	ID       int64
	Messages int64
}

func NewUser(id int64) *User {
	return &User{
		ID:       id,
		Messages: 0,
	}
}
