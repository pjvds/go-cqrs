package serialization

import (
	"github.com/dominikmayer/go-cqrs/storage"
	. "github.com/smartystreets/goconvey/convey"
	. "launchpad.net/gocheck"
	"reflect"
	"testing"
)

type AType struct{}
type BType struct{}
type CType struct{}

type EventTypeRegisterTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventTypeRegisterTestSuite{})

func (suite *EventTypeRegisterTestSuite) SetUpSuite(c *C) {
}

func TestRegistersAlwaysElementTypes(t *testing.T) {
	Convey("Given a register and three EventNames A, B and C of types AType, BType and CType", t, func() {
		register := NewEventTypeRegister()

		aName := *storage.NewEventName("A")
		bName := *storage.NewEventName("B")
		cName := *storage.NewEventName("C")

		Convey("We register A by instance", func() { //TODO: Is A or B the pointer?
			register.RegisterInstance(aName, AType{})

			Convey("When we get the type of A", func() {
				aType, _ := register.Get(aName)

				Convey("Then it should be of kind Struct", func() {
					So(aType.Kind(), ShouldEqual, reflect.Struct)
				})
			})
		})
		Convey("We register B by instance with a pointer type", func() { //TODO: Is A or B the pointer?
			register.RegisterInstance(bName, &BType{})

			Convey("When we get the type of B", func() {
				bType, _ := register.Get(aName)

				Convey("Then it should be of kind Struct", func() {
					So(bType.Kind(), ShouldEqual, reflect.Struct)
				})
			})
		})
		Convey("We register C by instance with a variable of pointer type", func() { //TODO: Is A or B the pointer?
			ctype := &CType{}
			register.RegisterInstance(cName, &ctype)

			Convey("When we get the type of C", func() {
				cType, _ := register.Get(aName)

				Convey("Then it should be of kind Struct", func() {
					So(cType.Kind(), ShouldEqual, reflect.Struct)
				})
			})
		})
	})
}

func (suite *EventTypeRegisterTestSuite) TestRegistersAlwaysElementTypesAlt(c *C) {
	aName := *storage.NewEventName("A")
	bName := *storage.NewEventName("B")
	cName := *storage.NewEventName("C")

	register := NewEventTypeRegister()
	register.RegisterInstance(aName, AType{})
	register.RegisterInstance(bName, &BType{})

	ctype := &CType{}
	register.RegisterInstance(cName, &ctype)

	aType, _ := register.Get(aName)
	bType, _ := register.Get(bName)
	cType, _ := register.Get(cName)

	c.Assert(aType.Kind(), Equals, reflect.Struct)
	c.Assert(bType.Kind(), Equals, reflect.Struct)
	c.Assert(cType.Kind(), Equals, reflect.Struct)
}

func (suite *EventTypeRegisterTestSuite) TestGetReturnsOkForKnownType(c *C) {
	aName := *storage.NewEventName("A")

	register := NewEventTypeRegister()
	register.RegisterInstance(aName, AType{})

	_, ok := register.Get(aName)
	c.Assert(ok, Equals, true)
}

func (suite *EventTypeRegisterTestSuite) TestGetReturnsNOkForUnknownType(c *C) {
	aName := *storage.NewEventName("A")

	register := NewEventTypeRegister()

	_, ok := register.Get(aName)
	c.Assert(ok, Equals, false)
}

func (suite *EventTypeRegisterTestSuite) TestGetReturnsType(c *C) {
	aName := *storage.NewEventName("A")

	register := NewEventTypeRegister()
	register.RegisterInstance(aName, AType{})

	t, _ := register.Get(aName)
	c.Assert(t, NotNil)
}
