package value_objects

import "github.com/google/uuid"

type ID struct {
	uuid.UUID
}

func NewID() *ID {
	return &ID{UUID: uuid.New()}
}

func (id *ID) ParseToString() string {
	return id.String()
}
