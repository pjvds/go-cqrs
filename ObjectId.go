package cqrs

import (
	"github.com/pjvds/gouuid"
)

type ObjectId uuid.UUID

// Creates a new unique ObjectId.
func NewObjectId() ObjectId {
	guid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return ObjectId(*guid)
}

// Creates a ObjectId from the given string value. It
// accepts strings in the following format: "6ba7b814-9dad-11d1-80b4-00c04fd430c8"
func ParseObjectId(value string) (id ObjectId, err error) {
	guid := new(uuid.UUID)
	if guid, err = uuid.ParseHex(value); err == nil {
		id = ObjectId(*guid)
	}

	return
}

// Returns a string representation of the ObjectId, like 6ba7b814-9dad-11d1-80b4-00c04fd430c8.
func (id ObjectId) String() string {
	guid := uuid.UUID(id)
	return guid.String()
}

// Returns the JSON encoding of the id.
func (id ObjectId) MarshalJSON() ([]byte, error) {
	value := uuid.UUID(id)
	return value.MarshalJSON()
}

// Sets id to the copy of the data.
func (id *ObjectId) UnmarshalJSON(b []byte) error {
	value := uuid.UUID(*id)
	err := value.UnmarshalJSON(b)
	if err != nil {
		return err
	}

	*id = ObjectId(value)
	return nil
}
