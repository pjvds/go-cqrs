package sourcing

import (
	"testing"
	"encoding/json"
	. "launchpad.net/gocheck"
)

// The state for the test suite
type EventSourceIdTestSuite struct {
}

func TestEventSourceId(t *testing.T) { TestingT(t) }

// Setup the test suite
//var _ = Suite(&EventSourceIdTestSuite{})

// Make sure we can turn an EventSourceId into a JSON value
func (s *EventSourceIdTestSuite) TestMarshallJSON(c *C) {
	id := NewEventSourceId()

	t := &struct {
		Id    EventSourceId  `json:"id"`
		IdPtr *EventSourceId `json:"idPtr"`
	}{
		Id:    id,
		IdPtr: &id,
	}

	b, err := json.Marshal(t)
	c.Assert(err, IsNil)

	c.Assert(string(b), Equals, "{\"id\":\""+t.Id.String()+"\",\"idPtr\":\""+t.IdPtr.String()+"\"}")
}

// Make sure we can turn an JSON value into an EventSourceId
func (s *EventSourceIdTestSuite) TestUnMarshallJSON(c *C) {
	t := &struct {
		Id    EventSourceId  `json:"id"`
		IdPtr *EventSourceId `json:"idPtr"`
	}{}

	id := NewEventSourceId()
	data := []byte("{\"id\":\"" + id.String() + "\",\"idPtr\":\"" + id.String() + "\"}")
	err := json.Unmarshal(data, &t)
	c.Assert(err, IsNil)

	c.Assert(t.Id, Equals, id)
	c.Assert(*t.IdPtr, Equals, id)
}
