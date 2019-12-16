package sourcing

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Make sure we can turn an EventSourceId into a JSON value
func TestMarshallJSON(t *testing.T) {
	Convey("Given we have an object with an id and a pointer to that id", t, func() {
		id := NewEventSourceId()

		testobject := &struct {
			Id    EventSourceId  `json:"id"`
			IdPtr *EventSourceId `json:"idPtr"`
		}{
			Id:    id,
			IdPtr: &id,
		}

		Convey("When we marshall the object", func() {
			b, err := json.Marshal(testobject)

			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the object is marshalled correctly", func() {
				So(string(b), ShouldEqual, "{\"id\":\""+testobject.Id.String()+"\",\"idPtr\":\""+testobject.IdPtr.String()+"\"}")
			})

		})
	})
}

// Make sure we can turn an JSON value into an EventSourceId
func TestUnMarshallJSON(t *testing.T) {
	Convey("Given we have the JSON representation of an object with an id and a pointer to that id", t, func() {

		id := NewEventSourceId()

		testobject := &struct {
			Id    EventSourceId  `json:"id"`
			IdPtr *EventSourceId `json:"idPtr"`
		}{}

		data := []byte("{\"id\":\"" + id.String() + "\",\"idPtr\":\"" + id.String() + "\"}")

		Convey("When we unmarshall the object", func() {
			err := json.Unmarshal(data, &testobject)

			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The unmarshalled id is equal to the original one", func() {
				So(testobject.Id, ShouldResemble, id)
			})

			Convey("And the unmarshalled pointer points to the id", func() {
				So(*testobject.IdPtr, ShouldResemble, id)
				// So(testobject.IdPtr, ShouldPointTo, &testobject.Id) //TODO
			})
		})
	})
}
