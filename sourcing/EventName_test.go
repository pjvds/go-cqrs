package sourcing

import (
	"encoding/json"
	. "launchpad.net/gocheck"
)

// The state for the test suite
type EventNameTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventNameTestSuite{})

// Make sure we can turn an EventName into a JSON value
func (s *EventNameTestSuite) TestMarshallJSON(c *C) {
	t := &struct {
		Name    EventName  `json:"name"`
		NamePtr *EventName `json:"namePtr"`
	}{
		Name:    *NewEventName("Name value"),
		NamePtr: NewEventName("NamePtr value"),
	}

	b, err := json.Marshal(t)
	c.Assert(err, IsNil)

	c.Assert(string(b), Equals, "{\"name\":\"Name value\",\"namePtr\":\"NamePtr value\"}")
}

// Make sure we can turn an JSON value into an EventName
func (s *EventNameTestSuite) TestUnMarshallJSON(c *C) {
	t := &struct {
		Name    EventName  `json:"name"`
		NamePtr *EventName `json:"namePtr"`
	}{}

	data := []byte("{\"name\":\"Name value\",\"namePtr\":\"NamePtr value\"}")
	err := json.Unmarshal(data, &t)
	c.Assert(err, IsNil)

	c.Assert(t.Name, Equals, *NewEventName("Name value"))
	c.Assert(*t.NamePtr, Equals, *NewEventName("NamePtr value"))
}
