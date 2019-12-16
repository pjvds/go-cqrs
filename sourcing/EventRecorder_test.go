package sourcing

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEventRecorder(t *testing.T) {
	Convey("Given an event and a recorder", t, func() {
		event := Event(struct {
			Foo string
		}{
			Foo: "bar",
		})

		recorder := NewEventRecorder()

		Convey("When we record the event", func() {
			recorder.Record(&event)

			Convey("Then the recorder should have one event recorded", func() {
				So(len(recorder.GetEvents()), ShouldEqual, 1)

				Convey("And the event should resemble the original one", func() {
					recordedEvents := recorder.GetEvents()
					So(recordedEvents[0], ShouldResemble, &event) // TODO: Why is it a pointer??
				})
			})
		})
		Convey("When we record the event again", func() {
			recorder.Record(&event)

			Convey("Then the recorder should have two events recorded", func() {
				So(len(recorder.GetEvents()), ShouldEqual, 2)

				Convey("And the second event should also resemble the original one", func() {
					recordedEvents := recorder.GetEvents()
					So(recordedEvents[1], ShouldResemble, &event)
				})
			})
			Convey("When we clear the recorder", func() {
				recorder.Clear()

				Convey("Then the recorder should have no more events recorded", func() {
					So(len(recorder.GetEvents()), ShouldEqual, 0)

					Convey("And the recorder should be able to record the event again", func() {
						recorder.Record(&event)
						So(len(recorder.GetEvents()), ShouldEqual, 1)
					})
				})
			})
		})

	})
}
