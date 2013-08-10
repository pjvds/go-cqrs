package cqrs

import (
	"github.com/pjvds/gouuid"
)

type ObjectId uuid.UUID

func NewObjectId() ObjectId {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return ObjectId(*guid)
}

func ParseObjectId(value string) (id ObjectId, err error) {
	guid := new(uuid.UUID)
	if guid, err = uuid.ParseHex(value); err == nil {
		id = ObjectId(*guid)
	}

	return
}

func (id ObjectId) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}

func (id ObjectId) MarshalJSON() ([]byte, error) {
	value := uuid.UUID(id)
	return value.MarshalJSON()
}

func (id *ObjectId) UnmarshalJSON(b []byte) error {
	value := uuid.UUID(*id)
	err := value.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	*id = ObjectId(value)
	return nil
}
