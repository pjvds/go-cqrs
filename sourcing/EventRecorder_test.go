package sourcing

import (
	"testing"
	. "launchpad.net/gocheck"
)

// The state for the test suite
type EventRecorderTestSuite struct {
}

func TestEventRecorder(t *testing.T) { TestingT(t) }

// Setup the test suite
//var _ = Suite(&EventRecorderTestSuite{})

// Make sure we can record events
func (s *EventRecorderTestSuite) TestRecord(c *C) {
	event := Event(struct {
		Foo string
	}{
		Foo: "bar",
	})

	recoder := NewEventRecorder()
	recoder.Record(&event)

	c.Assert(recoder.GetEvents(), HasLen, 1)
}

// Make sure we can record events, even if they are the same
func (s *EventRecorderTestSuite) TestRecordSameEventTwice(c *C) {
	event := Event(struct {
		Foo string
	}{
		Foo: "bar",
	})

	recoder := NewEventRecorder()
	recoder.Record(&event)
	recoder.Record(&event)

	c.Assert(recoder.GetEvents(), HasLen, 2)
}

func (s *EventRecorderTestSuite) TestClearClears(c *C) {
	event := Event(struct {
		Foo string
	}{
		Foo: "bar",
	})

	recoder := NewEventRecorder()
	recoder.Record(&event)

	c.Assert(recoder.GetEvents(), HasLen, 1)

	recoder.Clear()

	c.Assert(recoder.GetEvents(), HasLen, 0)
}

func (s *EventRecorderTestSuite) TestCanRecordAfterClear(c *C) {
	event := Event(struct {
		Foo string
	}{
		Foo: "bar",
	})

	recoder := NewEventRecorder()
	recoder.Record(&event)
	recoder.Clear()
	recoder.Record(&event)

	c.Assert(recoder.GetEvents(), HasLen, 1)
}
