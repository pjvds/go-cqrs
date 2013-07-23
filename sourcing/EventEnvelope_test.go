package sourcing

import (
	. "launchpad.net/gocheck"
	"time"
)

// The state for the test suite
type EventEnvelopeTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventEnvelopeTestSuite{})

func (s *EventEnvelopeTestSuite) TestEventEnvelopeString(c *C) {
	e := NewEventEnvelope(NewEventSourceId(), NewEventId(), "eventname", 0, time.Now(), "payload")
	c.Assert(e.String(), Equals, "0/eventname")
}

func (s *EventEnvelopeTestSuite) TestNewEventEnvelope(c *C) {
	eventNamer := NewTypeEventNamer()
	payload := struct {
		Foo string
		Bar int
	}{
		Foo: "baz",
		Bar: 42,
	}

	eventSourceId := NewEventSourceId()
	eventId := NewEventId()
	name := eventNamer.GetEventName(payload)
	sequence := NewEventSequence(0)
	timestamp := time.Now()

	sut := NewEventEnvelope(eventSourceId, eventId, name, sequence, timestamp, payload)

	c.Assert(sut.EventSourceId, Equals, eventSourceId)
	c.Assert(sut.EventId, Equals, eventId)
	c.Assert(sut.Name, Equals, name)
	c.Assert(sut.Sequence, Equals, sequence)
	c.Assert(sut.Timestamp, Equals, timestamp)
	c.Assert(sut.Payload, Equals, payload)
}
