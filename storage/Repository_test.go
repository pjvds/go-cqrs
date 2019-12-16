package storage

import (
//	"github.com/dominikmayer/go-cqrs/sourcing"
//	"github.com/dominikmayer/go-cqrs/storage/memory"
	. "launchpad.net/gocheck"
	"fmt"
)

// The state for the test suite
type RepositoryTestSuite struct {
}

//type TestInterceptor struct {
//	TotalDispatchCount int
//	LastDispatched     *EventStreamChange
//}
//
//func (t *TestInterceptor) Dispatch(change *EventStreamChange) {
//	t.LastDispatched = change
//	t.TotalDispatchCount++
//}

//Setup the test suite
var _ = Suite(&RepositoryTestSuite{})

func (s *RepositoryTestSuite) TestDispatchGetsCalledAfterAdd(c *C) {
	//result := NewRepository(memory.NewMemoryBackend())
	fmt.Printf("Hello world 2!")
}
