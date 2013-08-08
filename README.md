# Go-CQRS

[![wercker status](https://app.wercker.com/status/2b9662864982402105e0e2555a8a20da "wercker status")](https://app.wercker.com/project/bykey/2b9662864982402105e0e2555a8a20da)

This is an experimental library that tries to bring the concepts from CQRS to Go. It
currently focusses on adding event sourcing. Event sourcing ensures that all changes
to the application state are stored as a sequence of events. Not just can we query
these events, we can also use these events to reconstruct past and current state.

## Event sourcing example

Here are the examples of the main concepts.

### Changes to an object are captured by events

``` go
// Create a new domain object
user := domain.NewUser("pjvds")
c.Assert(user.Username, Equals, "pjvds")

// We created a new user, this should be
// captured by an event.
c.Assert(len(user.Events()), Equals, 1)

// Change the username of the user
user.ChangeUsername("wwwouter")
c.Assert(user.Username, Equals, "wwwouter")

// We changed the username, this should be
// captured by an event.
c.Assert(len(user.Events()), Equals, 2)
```

### An object state can be rebuild from history

``` go
// The id of our event source that we will rebuild from history.
sourceId, _ := sourcing.ParseEventSourceId("0791d279-664d-458e-bf60-567ade140832")

// The full history for the User domain object
history := []sourcing.Event{
    // It was first created
    events.UserCreated{
        Username: "pjvds",
    },
    // Then the username was changed
    events.UsernameChanged{
        OldUsername: "pjvds",
        NewUsername: "wwwouter",
    },
}

// Create a new User domain object from history
user := domain.NewUserFromHistory(sourceId, history)

// It should not have the initial state.
c.Assert(user.Username, Not(Equals), "pjvds")

// It should have the latest state.
c.Assert(user.Username, Equals, "wwwouter")
```

## Sourcing object

An sourcing object has three main concepts: state, command methods and event handlers.

### State

An objects starts with a struct that holds the state of an object. In this case
a simple `User` object that has a single field `Username`. Not the comment that
we will never update this username directly.

``` go
// Holds the state of our user. Note that
// state like Username is not updated directly!
type User struct {
    // Embed event source funtionality.
    sourcing.EventSource

    Username string
}
```

### Ctor

The `User` has two constructor methods. One to create a new `User` and one that
creates a `User` based on history. The first creates a `User` and applies the fact
that this happend. The later creates a `User` and replays history to build the state.

``` go
// Creates an new User object.
func NewUser(username string) *User {
    // Create a new user object
    user := new(User)
    user.EventSource = sourcing.CreateNew(user)

    // Apply the fact that the user was created.
    user.Apply(events.UserCreated{
        Username: username,
    })

    return user
}

// Creates an new User object and builds the state from the history.
func NewUserFromHistory(history []sourcing.EventEnvelope) *User {
    var user = new(User)
    user.EventSource = sourcing.AttachFromHistory(user, history)

    return user
}
```

### Command method

The `User` has only one exported fields called `Username`. The state of an `User`
object, or any other sourced object, should never be updated directly. Rather
they are updated via methods that capture the intention of what should be done.
These methods are called command methods. These methods validate whether the
state change can happen, and if so, they create an event that represents this
change and the event is applied. This means the event is recorded and send to
the event handler on the object itself where it will update it's internal state.

``` go
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
```

### Event handler

The `user.Apply()` call registers that the event is happened and calls
an event handling method on the `User` object that updates the state occordingly.

``` go
// Update the User state for an UserCreated event.
func (user *User) HandleUserCreated(e events.UserCreated) {
    user.Username = e.Username
}

// Update the User state for an UsernameChanged event.
func (user *User) HandleUsernameChanged(e events.UsernameChanged) {
    user.Username = e.NewUsername
}
```

See the [User.go](https://github.com/pjvds/go-cqrs/blob/master/tests/domain/User.go)
source to see all the details.
