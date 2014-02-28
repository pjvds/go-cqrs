package storage

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestStruct struct {
	Name    EventName  `json:"name"`
	NamePtr *EventName `json:"namePtr"`
}

var nameValue string = "Name value"

var namePtrValue string = "NamePtr value"

var testEventName EventName = *NewEventName(nameValue)

var testEventNamePtr *EventName = NewEventName(namePtrValue)

var testobjectJson string = "{\"name\":\"" + nameValue + "\",\"namePtr\":\"" + namePtrValue + "\"}"

func TestMarshalEventName(t *testing.T) {
	Convey("Given we have an object with an EventName and a pointer to it and the correct string representation", t, func() {

		testobject := TestStruct{
			Name:    testEventName,
			NamePtr: testEventNamePtr,
		}

		Convey("When we marshal the object", func() {
			b, err := json.Marshal(testobject)

			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the string matches the correct string representation", func() {
				So(string(b), ShouldEqual, testobjectJson)
			})
		})
	})
}

// Make sure we can turn an JSON value into an EventName
func TestUnmarshalEventName(t *testing.T) {
	Convey("Given we have the string representation of an object with an EventName and a pointer to it", t, func() {

		testobject := TestStruct{}

		data := []byte(testobjectJson)

		Convey("When we unmarshal the string representation", func() {

			err := json.Unmarshal(data, &testobject)

			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The EventName matches the original one", func() {
				So(testobject.Name, ShouldResemble, testEventName)
			})

			Convey("And the pointer to EventName matches the original one", func() {
				So(testobject.NamePtr, ShouldResemble, testEventNamePtr)
			})
		})
	})
}
