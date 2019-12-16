package serialization

import (
	"github.com/dominikmayer/go-cqrs/storage"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type AType struct{}
type BType struct{}
type CType struct{}

func TestEventTypeRegister(t *testing.T) {
	Convey("Given a register and three EventNames A, B and C of types AType, BType and CType", t, func() {
		register := NewEventTypeRegister()

		aName := *storage.NewEventName("A")
		bName := *storage.NewEventName("B")
		cName := *storage.NewEventName("C")

		Convey("When we try to get the type of A", func() {
			_, notOk := register.Get(aName)

			Convey("Then Get should return 'not ok'/false", func() {
				So(notOk, ShouldBeFalse)
			})
		})

		Convey("We register A by instance", func() { //TODO: Is A or B the pointer?
			register.RegisterInstance(aName, AType{})
		})
		Convey("We register B by instance with a pointer type", func() { //TODO: Is A or B the pointer?
			register.RegisterInstance(bName, &BType{})
		})
		Convey("We register C by instance with a variable of pointer type", func() { //TODO: Is A or B the pointer?
			ctype := &CType{}
			register.RegisterInstance(cName, &ctype)
		})
		Convey("When we get the types", func() {
			aType, okA := register.Get(aName)
			bType, okB := register.Get(bName)
			cType, okC := register.Get(cName)

			Convey("Then they should not be nil", func() {
				So(aType, ShouldNotBeNil)
				So(bType, ShouldNotBeNil)
				So(cType, ShouldNotBeNil)
			})

			Convey("They should all be of kind Struct", func() {
				So(aType.Kind(), ShouldEqual, reflect.Struct)
				So(bType.Kind(), ShouldEqual, reflect.Struct)
				So(cType.Kind(), ShouldEqual, reflect.Struct)
			})

			Convey("And Get should have returned 'ok'", func() {
				So(okA, ShouldBeTrue)
				So(okB, ShouldBeTrue)
				So(okC, ShouldBeTrue)
			})
		})
	})
}
