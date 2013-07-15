package domain

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/events"
)

// Holds the state of our user. Note that
// state like Username is not updated directly!
type User struct {
	sourcer sourcing.EventSource

	Username string
}

// Creates an new User object.
func NewUser(username string) *User {
	user := new(User)
	user.sourcer = sourcing.AttachNew(user)

	user.sourcer.Apply(events.UserCreated{
		Username: username,
	})

	return user
}

// Creates an new User object and builds the state from the history.
func NewUserFromHistory(history []sourcing.EventEnvelope) *User {
	var user = new(User)
	user.sourcer = sourcing.AttachFromHistory(user, history)

	return user
}

// Change the username to a new name.
func (user *User) ChangeUsername(username string) {
	user.sourcer.Apply(events.UsernameChanged{
		OldUsername: user.Username,
		NewUsername: username,
	})
}

// Update the User state for an UserCreated event.
func (user *User) HandleUserCreated(e events.UserCreated) {
	user.Username = e.Username
}

// Update the User state for an UsernameChanged event.
func (user *User) HandleUsernameChanged(e events.UsernameChanged) {
	user.Username = e.NewUsername
}
