package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomUUID(t *testing.T) {
	id := GenerateRandomUUID()

	assert.IsType(t, uuid.UUID{}, id)
}
