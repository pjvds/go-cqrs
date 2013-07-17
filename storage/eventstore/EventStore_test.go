package eventstore

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/domain"
	//"github.com/pjvds/go-cqrs/tests/events"
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	InitLogging()
	TestingT(t)
}

// The state for the test suite
type EventStoreTestSuite struct {
	store *EventStore
}

// Setup the test suite
var _ = Suite(&EventStoreTestSuite{})

func (s *EventStoreTestSuite) SetUpSuite(c *C) {
	store, _ := DailEventStore("http://127.0.0.1:2113")
	s.store = store
}

func (s *EventStoreTestSuite) TestSmoke(c *C) {
	// Create a new domain object
	user := domain.NewUser("pjvds")
	state := sourcing.GetState(user)
	s.store.NewStream(state)
}
