package storage

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type myEvent struct {
	Foo string
	Bar int
}

func TestEventNamer(t *testing.T) {
	Convey("When we request a new TypeEventNamer", t, func() {
		result := NewTypeEventNamer()

		Convey("Then we should be given one", func() {
			So(result, ShouldNotBeNil)
			So(result, ShouldHaveSameTypeAs, &TypeEventNamer{})
		})

		Convey("When we have an event 'myEvent' in the package 'github.com/pjvds/go-cqrs/storage'", func() {
			event := myEvent{
				Foo: "baz",
				Bar: 42,
			}
			Convey("And we request the event name", func() {
				name := result.GetEventName(event)

				Convey("Then it should be 'github.com/pjvds/go-cqrs/storage/myEvent'", func() {
					So(name, ShouldEqual, "github.com/pjvds/go-cqrs/storage/myEvent")
				})
			})
		})

		Convey("When we have a pointer to 'myEvent' in the package 'github.com/pjvds/go-cqrs/storage'", func() {
			event := &myEvent{
				Foo: "baz",
				Bar: 42,
			}
			Convey("And we request the event name", func() {
				name := result.GetEventName(event)

				Convey("Then it should be 'github.com/pjvds/go-cqrs/storage/myEvent'", func() {
					So(name, ShouldEqual, "github.com/pjvds/go-cqrs/storage/myEvent")
				})
			})
		})
	})
}
