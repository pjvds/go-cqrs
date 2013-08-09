package storage

import (
	"encoding/json"
	. "launchpad.net/gocheck"
)

// The state for the test suite
type EventStreamIdTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventStreamIdTestSuite{})

// Make sure we can turn an EventStreamId into a JSON value
func (s *EventStreamIdTestSuite) TestMarshallJSON(c *C) {
	id := NewEventStreamId()

	t := &struct {
		Id    EventStreamId  `json:"id"`
		IdPtr *EventStreamId `json:"idPtr"`
	}{
		Id:    id,
		IdPtr: &id,
	}

	b, err := json.Marshal(t)
	c.Assert(err, IsNil)

	c.Assert(string(b), Equals, "{\"id\":\""+t.Id.String()+"\",\"idPtr\":\""+t.IdPtr.String()+"\"}")
}

// Make sure we can turn an JSON value into an EventStreamId
func (s *EventStreamIdTestSuite) TestUnMarshallJSON(c *C) {
	t := &struct {
		Id    EventStreamId  `json:"id"`
		IdPtr *EventStreamId `json:"idPtr"`
	}{}

	id := NewEventStreamId()
	data := []byte("{\"id\":\"" + id.String() + "\",\"idPtr\":\"" + id.String() + "\"}")
	err := json.Unmarshal(data, &t)
	c.Assert(err, IsNil)

	c.Assert(t.Id, Equals, id)
	c.Assert(*t.IdPtr, Equals, id)
}
