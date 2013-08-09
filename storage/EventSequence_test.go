package storage

import (
	. "launchpad.net/gocheck"
)

// The state for the test suite
type EventSequenceTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventSequenceTestSuite{})

func (s *EventSequenceTestSuite) TestTwoSequencesWithSameValueAreTheSame(c *C) {
	sequenceA := NewEventSequence(0)
	sequenceB := NewEventSequence(0)

	c.Assert(sequenceA, Equals, sequenceB)
}

func (s *EventSequenceTestSuite) TestSequenceNextAddsOne(c *C) {
	sequence := NewEventSequence(0)
	expectedSequence := NewEventSequence(1)

	actual := sequence.Next()
	c.Assert(actual, Equals, expectedSequence)
}
