package domain

import (
	"errors"

	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/events"
)

// Holds the state of our user. Note that the
// state, like Username, is not updated directly.
type User struct {
	// Make this user an event source.
	// This fields needs to be initialized at construction before any state is changed.
	sourcing.EventSource

	// Holds the username, do not update this fields directly.
	// It should only be modified by ctor and ChangeUsername method.
	Username string
}

// Creates an new User object.
func NewUser(username string) *User {
	user := new(User)
	user.EventSource = sourcing.CreateNew(user)

	user.Apply(events.UserCreated{
		Username: username,
	})

	return user
}

// Creates an new User object and builds the state from the history.
func NewUserFromHistory(sourceId sourcing.EventSourceId, history []sourcing.Event) *User {
	var user = new(User)
	user.EventSource = sourcing.CreateFromHistory(user, sourceId, history)

	return user
}

// Change the username to a new name.
func (user *User) ChangeUsername(username string) error {
	// Validate username
	if length := len(username); length < 3 || length > 20 {
		return errors.New("invalid username length")
	}

	// Raise the fact that the username is changed.
	user.Apply(events.UsernameChanged{
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
