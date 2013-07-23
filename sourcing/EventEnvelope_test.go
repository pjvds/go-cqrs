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
