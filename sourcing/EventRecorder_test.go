package sourcing

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Make sure we can record events
func TestRecord(t *testing.T) {
	Convey("Given an event and a recorder", t, func() {
		event := Event(struct {
			Foo string
		}{
			Foo: "bar",
		})

		recorder := NewEventRecorder()

		Convey("When we record the event", func() {
			recorder.Record(&event)
		})

		Convey("Then the recorder should have one event recorded", func() {
			So(len(recorder.GetEvents()), ShouldEqual, 1)
		})

		Convey("And the event should resemble the recorded one", func() {
			recordedEvents := recorder.GetEvents()
			So(recordedEvents[0], ShouldResemble, &event)
		})
	})
}

// Make sure we can record events, even if they are the same
func TestRecordSameEventTwice(t *testing.T) {
	Convey("Given an event and a recorder", t, func() {
		event := Event(struct {
			Foo string
		}{
			Foo: "bar",
		})

		recorder := NewEventRecorder()

		Convey("When we record the event twice", func() {
			recorder.Record(&event)
			recorder.Record(&event)
		})

		Convey("Then the recorder should have two events recorded", func() {
			So(len(recorder.GetEvents()), ShouldEqual, 2)
		})

		Convey("And both events should resemble the recorded one", func() {
			recordedEvents := recorder.GetEvents()
			So(recordedEvents[0], ShouldResemble, &event)
			So(recordedEvents[1], ShouldResemble, &event)
		})
	})
}

func TestClearClears(t *testing.T) {
	Convey("Given an event and a recorder", t, func() {
		event := Event(struct {
			Foo string
		}{
			Foo: "bar",
		})

		recorder := NewEventRecorder()

		Convey("When we record the event", func() {
			recorder.Record(&event)
			So(len(recorder.GetEvents()), ShouldEqual, 1)
		})

		Convey("And clear the recorder", func() {
			recorder.Clear()
		})

		Convey("Then the recorder should have no more events recorded", func() {
			So(len(recorder.GetEvents()), ShouldEqual, 0)
		})
	})
}

func TestCanRecordAfterClear(t *testing.T) {
	Convey("Given an event and a recorder", t, func() {
		event := Event(struct {
			Foo string
		}{
			Foo: "bar",
		})

		recorder := NewEventRecorder()

		Convey("When we record the event", func() {
			recorder.Record(&event)
			So(len(recorder.GetEvents()), ShouldEqual, 1)
		})

		Convey("And clear the recorder", func() {
			recorder.Clear()
			So(len(recorder.GetEvents()), ShouldEqual, 0)
		})

		Convey("Then the recorder should be able to record the event again", func() {
			recorder.Record(&event)
			So(len(recorder.GetEvents()), ShouldEqual, 1)
		})
	})
}
