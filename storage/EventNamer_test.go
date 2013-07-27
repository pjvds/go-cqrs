package storage

import (
	. "launchpad.net/gocheck"
)

// The state for the test suite
type EventNamerTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventNamerTestSuite{})

type myEvent struct {
	Foo string
	Bar int
}

func (s *EventNamerTestSuite) TestNewTypeEventNamerReturnsValue(c *C) {
	result := NewTypeEventNamer()
	c.Assert(result, NotNil)
}

func (s *EventNamerTestSuite) TestGetEventName(c *C) {
	event := myEvent{
		Foo: "baz",
		Bar: 42,
	}

	result := NewTypeEventNamer()
	name := result.GetEventName(event)

	c.Assert(name, Equals, EventName("github.com/pjvds/go-cqrs/sourcing/myEvent"))
}
