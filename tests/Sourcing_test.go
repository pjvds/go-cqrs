package tests

import (
	"github.com/pjvds/go-cqrs/sourcing"
	"github.com/pjvds/go-cqrs/tests/domain"
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	InitLogging()
	TestingT(t)
}

// The state for the test suite
type AppTestSuite struct {
}

// Setup the test suite
var _ = Suite(&AppTestSuite{})

func (s *AppTestSuite) TestStateIsEventSources(c *C) {
	// Create a new domain object
	user := domain.NewUser("pjvds")
	c.Assert(user.Username, Equals, "pjvds")

	// We created a new user, this should be
	// captured by an event.
	state := sourcing.GetState(user)
	c.Assert(len(state.Events), Equals, 1)

	// Change the username of the user
	user.ChangeUsername("wwwouter")
	c.Assert(user.Username, Equals, "wwwouter")

	// We changed the username, this should be
	// captured by an event.
	c.Assert(len(state.Events), Equals, 2)
}
