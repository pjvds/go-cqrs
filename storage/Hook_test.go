package storage

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

// The state for the test suite
type EventNamerTestSuite struct {
}

// Setup the test suite
var _ = Suite(&EventNamerTestSuite{})
