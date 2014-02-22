package serialization

import (
	"github.com/dominikmayer/go-cqrs/storage"
)

type Serializer interface {
	Serialize(e *storage.Event) ([]byte, error)
	Deserialize(name storage.EventName, data []byte) (*storage.Event, error)
}
