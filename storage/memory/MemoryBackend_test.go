package memory

import (
	"testing"
	"encoding/json"
	. "github.com/dominikmayer/go-cqrs/storage"
	. "launchpad.net/gocheck"
	"time"
)

type myEvent struct {
	Foo string
	Bar int
}

// The state for the test suite
type MemoryBackendSuite struct {
}

func TestMemoryBackend(t *testing.T) { TestingT(t) }

// Setup the test suite
var _ = Suite(&MemoryBackendSuite{})

func (suite *MemoryBackendSuite) TestWriteStreamDoesNotError(c *C) {
	backend := NewMemoryBackend()
	change := createSomeChanges()

	err := backend.WriteStream(change)
	c.Assert(err, IsNil)
}

func createSomeChanges() *EventStreamChange {
	events := make([]*Event, 20)

	for i := 0; i < len(events); i++ {
		data, _ := json.Marshal(&myEvent{
			Foo: "foo",
			Bar: i,
		})

		events[i] = NewEvent(NewEventId(), *NewEventName("MyEvent"), NewEventSequence(int64(i)), time.Now(), data)
	}

	return &EventStreamChange{
		StreamId: NewEventStreamId(),
		From:     0,
		To:       int64(len(events)),
		Events:   events,
	}
}
