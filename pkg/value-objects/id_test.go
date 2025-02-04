package value_objects_test

import (
	value_objects "github.com/SchunckLeonardo/library-management-system/pkg/value-objects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	id := value_objects.NewID()
	assert.IsType(t, id, &value_objects.ID{})

	err := uuid.Validate(id.ParseToString())
	assert.Nil(t, err)
}

func TestID_ParseToString(t *testing.T) {
	id := value_objects.NewID()
	assert.Equal(t, id.ParseToString(), id.String())
}
