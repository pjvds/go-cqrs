package domain

import (
	"errors"
	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/events"
)

// Holds the state of our user. Note that
// state like Username is not updated directly!
type User struct {
	// Reference to the sourcer
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
func NewUserFromHistory(sourceId sourcing.EventSourceId, history []sourcing.Event) *User {
	var user = new(User)
	user.sourcer = sourcing.AttachFromHistory(user, sourceId, history)

	return user
}

// Disposes the user object and detach user from event sourcing context.
func (user *User) Dispose() {
	sourcing.Detach(user)
}

// Change the username to a new name.
func (user *User) ChangeUsername(username string) error {
	// Validate username
	if lenght := len(username); lenght < 3 || lenght > 20 {
		return errors.New("invalid username lenght")
	}

	// Raise the fact that the username is changed.
	user.sourcer.Apply(events.UsernameChanged{
		OldUsername: user.Username,
		NewUsername: username,
	})

	return nil
}

// Update the User state for an UserCreated event.
func (user *User) HandleUserCreated(e events.UserCreated) {
	user.Username = e.Username
}

// Update the User state for an UsernameChanged event.
func (user *User) HandleUsernameChanged(e events.UsernameChanged) {
	user.Username = e.NewUsername
}
