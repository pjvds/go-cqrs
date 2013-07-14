package domain

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/events"
)

type User struct {
	sourcer sourcing.EventSource

	Username string
}

func NewUser(username string) *User {
	user := new(User)
	user.sourcer = sourcing.AttachNew(user)

	user.sourcer.Apply(events.UserCreated{
		Username: username,
	})

	return user
}

func NewUserFromHistory(history []sourcing.EventEnvelope) *User {
	user := new(User)
	user.sourcer = sourcing.AttachFromHistory(user, history)

	return user
}

func (user *User) ChangeUsername(username string) {
	user.sourcer.Apply(events.UsernameChanged{
		OldUsername: user.Username,
		NewUsername: username,
	})
}

func (user *User) HandleUserCreated(e events.UserCreated) {
	user.Username = e.Username
}

func (user *User) HandleUsernameChanged(e events.UsernameChanged) {
	user.Username = e.NewUsername
}
