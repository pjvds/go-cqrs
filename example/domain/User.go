package domain

import (
	"github.com/pjvds/go-cqrs/eventing"
	"github.com/pjvds/go-cqrs/example/events"
)

type User struct {
	Username string
	applier  eventing.EventHandler
}

func NewUser(username string) *User {
	user := &User{}
	eventing.Attach(user)

	user.applier(events.UserCreated{
		Username: username,
	})

	return user
}

func (user *User) SetEventApplier(applier eventing.EventHandler) {
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
