package mongodb

import (
	"testing"
	//. "github.com/smartystreets/goconvey/convey"

	"flag"
	"fmt"
	"github.com/dominikmayer/go-cqrs/storage"
	"github.com/dominikmayer/go-cqrs/storage/serialization"
	"github.com/dominikmayer/go-cqrs/tests/domain"
	. "github.com/dominikmayer/go-cqrs/tests/events"
	. "launchpad.net/gocheck"
	"reflect"
)

type myEvent struct {
	Foo string
	Bar int
}

var testMongoDB = flag.Bool("mongodb", true, "Include MongoDB tests")

func init() {
	flag.Parse()
}

// The state for the test suite
type MongoDBTestSuite struct {
	store      *MongoDB
	repository *storage.Repository
}

func TestMongoDB(t *testing.T) { TestingT(t) }

// Setup the test suite
var _ = Suite(&MongoDBTestSuite{})

func (s *MongoDBTestSuite) SetUpSuite(c *C) {
	if !*testMongoDB {
		c.Skip("-MongoDB not provided")
	}

	register := serialization.NewEventTypeRegister()
	namer := storage.NewTypeEventNamer()

	userCreatedType := reflect.TypeOf(UserCreated{})
	userCreatedName := namer.GetEventNameFromType(userCreatedType)
	register.Register(userCreatedName, userCreatedType)

	usernameChangedType := reflect.TypeOf(UsernameChanged{})
	usernameChangedName := namer.GetEventNameFromType(usernameChangedType)
	register.Register(usernameChangedName, usernameChangedType)

	store := New("localhost:27017", "test", "user", register)
	s.store = store

	s.repository = storage.NewRepository(s.store, storage.NewNullEventDispatcher())
}

func (s *MongoDBTestSuite) TestSmoke(c *C) {
	// Create a new domain object
	toStore := domain.NewUser("pjvds")
	for i := 0; i < 24; i++ {
		toStore.ChangeUsername(fmt.Sprintf("pjvds%v", i))
	}

	err := s.repository.Add(toStore)
	c.Assert(err, IsNil)

	events, err := s.store.ReadStream(storage.EventStreamId(toStore.Id()))
	c.Assert(err, IsNil)
	c.Assert(len(events), Equals, 25)

	namer := storage.NewTypeEventNamer()
	userCreatedType := reflect.TypeOf(UserCreated{})
	c.Assert(events[0].Name, Equals, namer.GetEventNameFromType(userCreatedType))
}
