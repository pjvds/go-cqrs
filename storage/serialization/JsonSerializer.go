package serialization

import (
	"encoding/json"
	"github.com/pjvds/go-cqrs/sourcing"
)

type JsonSerializer struct {
}

func NewJsonSerializer() *JsonSerializer {
	return &JsonSerializer{}
}

func (s *JsonSerializer) Serialize(e *sourcing.Event) ([]byte, error) {
	data, err := json.Marshal(e)

	return data, err
}

func (s *JsonSerializer) Deserialize(data []byte, e *sourcing.Event) error {
	return json.Unmarshal(data, e)
}
