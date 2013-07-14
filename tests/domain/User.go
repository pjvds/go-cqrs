package domain

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/events"
)

type User struct {
	Username string
	applier  sourcing.EventHandler
}

func NewUser(username string) *User {
	user := new(User)
	sourcing.AttachNew(user)

	user.applier(events.UserCreated{
		Username: username,
	})

	return user
}

func (user *User) SetEventApplier(applier sourcing.EventHandler) {
	user.applier = applier
}

func (user *User) ChangeUsername(username string) {
	user.applier(events.UsernameChanged{
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
