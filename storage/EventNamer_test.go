package storage

import (
	. "launchpad.net/gocheck"
)

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

	c.Assert(name, Equals, EventName("github.com/pjvds/go-cqrs/storage/myEvent"))
}

func (s *EventNamerTestSuite) TestGetEventNameForPointer(c *C) {
	event := &myEvent{
		Foo: "baz",
		Bar: 42,
	}

	result := NewTypeEventNamer()
	name := result.GetEventName(event)

	c.Assert(name, Equals, EventName("github.com/pjvds/go-cqrs/storage/myEvent"))
}
