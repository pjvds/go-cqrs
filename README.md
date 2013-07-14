# Go-CQRS

This is an experimental library that tries to bring the concepts from CQRS to Go.

Currently I am trying to find a nice API to add event sourcing to object state.

## Event sourcing

sourcing sourcing ensures that all changes to the application state are stored
as a sequence of events. Not just can we query these events, we can also use
these events to reconstruct past and current state.

## How to run the examaple

    $ go build ./...
    $ cd example
    $ ./example
    2013/07/14 12:57:26 Attaching *domain.User
    2013/07/14 12:57:26 Applying *events.UserCreated to *domain.User
    2013/07/14 12:57:26 Event 1: *events.UserCreated

    2013/07/14 12:57:26 Bye!!
