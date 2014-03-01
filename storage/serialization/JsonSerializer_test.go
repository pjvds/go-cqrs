package serialization

import (
	"github.com/dominikmayer/go-cqrs/storage"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
	"time"
)

type FooEvent struct {
	Foo string
	Bar int
}

func TestSerialize(t *testing.T) {
	namer := storage.NewTypeEventNamer()

	event := new(FooEvent)
	eventType := reflect.TypeOf(event)
	eventName := namer.GetEventNameFromType(eventType)

	types := NewEventTypeRegister()
	types.Register(eventName, eventType)

	serializer := NewJsonSerializer(types)

	Convey("Given an event and its JSON", t, func() {
		expectedJson := "{\"eventId\":\"25f7fdb6-5ef9-47b0-55a1-b9160ce37730\",\"name\":\"github.com/dominikmayer/go-cqrs/storage/serialization/FooEvent\",\"sequence\":0,\"timestamp\":\"2011-02-01T12:32:40-05:00\",\"payload\":{\"Foo\":\"\",\"Bar\":0}}"

		timestamp, _ := time.Parse(time.RFC3339, "2011-02-01T12:32:40-05:00")

		eventId, err := storage.ParseEventId("\"25f7fdb6-5ef9-47b0-55a1-b9160ce37730\"")
		So(err, ShouldBeNil)

		event := new(FooEvent)

		envelope := storage.NewEvent(*eventId, eventName, storage.NewEventSequence(0), timestamp, event)

		Convey("When we serialize the event", func() {
			data, err := serializer.Serialize(envelope)

			Convey("Then there should be no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the serialized JSON should match the expected one", func() {
				actualJson := string(data)
				So(actualJson, ShouldEqual, expectedJson)
			})
		})
	})

	Convey("Given", func() {
		json := "{\"eventId\":\"25f7fdb6-5ef9-47b0-55a1-b9160ce37730\",\"name\":\"github.com/dominikmayer/go-cqrs/storage/serialization/FooEvent\",\"sequence\":0,\"timestamp\":\"2011-02-01T12:32:40-05:00\",\"payload\":{\"Foo\":\"hello world\",\"Bar\":42}}"

		Convey("Given", func() {
			event, err := serializer.Deserialize(eventName, []byte(json))

			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The data was retrieved", func() {
				So(event, ShouldNotBeNil)
			})

			Convey("It has the correct type", func() {
				So(event.Data, ShouldHaveSameTypeAs, new(FooEvent))
			})

			Convey("And the content was correctly deserialized", func() {
				So(event.Data.(*FooEvent).Foo, ShouldEqual, "hello world")
				So(event.Data.(*FooEvent).Bar, ShouldEqual, 42)
			})
		})
	})
}
