package serialization

import (
	"github.com/pjvds/go-cqrs/sourcing"
)

type Serializer interface {
	Serialize(e *sourcing.Event) ([]byte, error)
	Deserialize(data []byte, e *sourcing.Event) error
}
