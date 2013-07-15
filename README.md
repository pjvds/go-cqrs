# Go-CQRS

This is an experimental library that tries to bring the concepts from CQRS to Go.

Currently I am trying to find a nice API to add event sourcing to object state.

## Event sourcing

sourcing sourcing ensures that all changes to the application state are stored
as a sequence of events. Not just can we query these events, we can also use
these events to reconstruct past and current state.

## Event sourcing example

Here are the examples of the main concepts.

### Changes to an object are captured by events

``` go
// Create a new domain object
user := domain.NewUser("pjvds")
c.Assert(user.Username, Equals, "pjvds")

// We created a new user, this should be
// captured by an event.
state := sourcing.GetState(user)
c.Assert(len(state.Events()), Equals, 1)

// Change the username of the user
user.ChangeUsername("wwwouter")
c.Assert(user.Username, Equals, "wwwouter")

// We changed the username, this should be
// captured by an event.
c.Assert(len(state.Events()), Equals, 2)
```

### An object state can be rebuild from history

``` go
// The full history for the User domain object
history := sourcing.PackEvents([]sourcing.Event{
    // It was first created
    events.UserCreated{
        Username: "pjvds",
    },
    // Then the username was changed
    events.UsernameChanged{
        OldUsername: "pjvds",
        NewUsername: "wwwouter",
    },
})

// Create a new User domain object from history
user := domain.NewUserFromHistory(history)

c.Assert(user.Username, Not(Equals), "pjvds")
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
    sourcer sourcing.EventSource

    Username string
}
```

### Ctor

The `User` has two constructor methods. One to create a new `User` and one that
creates a `User` based on historical events. The first creates a `User` and
raises the fact that this happend by applying an event. The later creates a
`User` and replays history to build the state.

In both contrustor methods the `User` object is attached to the `sourcing`
context and the sourcer object is stored inside the `User`.

``` go
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

The `user.sourcer.Apply()` call registers that the event is happened and calls
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
