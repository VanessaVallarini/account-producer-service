package utils

import (
	"github.com/google/uuid"
)

func GenerateRandomUUID() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return id
}
