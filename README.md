# Go-CQRS

This is an experimental library that tries to bring the concepts from CQRS to Go.

Currently I am trying to find a nice API to add event sourcing to object state.

## Event sourcing

sourcing sourcing ensures that all changes to the application state are stored
as a sequence of events. Not just can we query these events, we can also use
these events to reconstruct past and current state.

## Example

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
