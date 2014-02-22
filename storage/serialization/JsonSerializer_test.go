package serialization

import (
	"github.com/dominikmayer/go-cqrs/storage"
	. "launchpad.net/gocheck"
	"reflect"
	"time"
)

type FooEvent struct {
	Foo string
	Bar int
}

type JsonSerializerTestSuite struct {
	fooEventType reflect.Type
	fooEventName storage.EventName

	serializer *JsonSerializer
}

// Setup the test suite
var _ = Suite(&JsonSerializerTestSuite{})

func (suite *JsonSerializerTestSuite) SetUpSuite(c *C) {
	namer := storage.NewTypeEventNamer()

	event := new(FooEvent)
	eventType := reflect.TypeOf(event)
	eventName := namer.GetEventNameFromType(eventType)

	types := NewEventTypeRegister()
	types.Register(eventName, eventType)

	suite.fooEventName = eventName
	suite.fooEventType = eventType
	suite.serializer = NewJsonSerializer(types)
}

func (suite *JsonSerializerTestSuite) TestSerialize(c *C) {
	expectedJson := "{\"eventId\":\"25f7fdb6-5ef9-47b0-55a1-b9160ce37730\",\"name\":\"github.com/dominikmayer/go-cqrs/storage/serialization/FooEvent\",\"sequence\":0,\"timestamp\":\"2011-02-01T12:32:40-05:00\",\"payload\":{\"Foo\":\"\",\"Bar\":0}}"

	timestamp, _ := time.Parse(time.RFC3339, "2011-02-01T12:32:40-05:00")

	eventId, err := storage.ParseEventId("\"25f7fdb6-5ef9-47b0-55a1-b9160ce37730\"")
	c.Assert(err, IsNil)

	event := new(FooEvent)

	envelope := storage.NewEvent(*eventId, suite.fooEventName, storage.NewEventSequence(0), timestamp, event)
	data, err := suite.serializer.Serialize(envelope)

	c.Assert(err, IsNil)

	actualJson := string(data)
	c.Assert(actualJson, Equals, expectedJson)
}

func (suite *JsonSerializerTestSuite) TestDerialize(c *C) {
	json := "{\"eventId\":\"25f7fdb6-5ef9-47b0-55a1-b9160ce37730\",\"name\":\"github.com/dominikmayer/go-cqrs/storage/serialization/FooEvent\",\"sequence\":0,\"timestamp\":\"2011-02-01T12:32:40-05:00\",\"payload\":{\"Foo\":\"hello world\",\"Bar\":42}}"
	serializer := suite.serializer

	event, err := serializer.Deserialize(suite.fooEventName, []byte(json))
	c.Assert(err, IsNil)
	c.Assert(event, NotNil)
	c.Assert(event.Data, FitsTypeOf, new(FooEvent))

	c.Assert(event.Data.(*FooEvent).Foo, Equals, "hello world")
	c.Assert(event.Data.(*FooEvent).Bar, Equals, 42)
}
