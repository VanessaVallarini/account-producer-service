package db

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewScylla(t *testing.T) {
	mockScylla := mocks.NewIScylla(t)

	assert.NotNil(t, mockScylla)
}

func TestScanMap(t *testing.T) {
	config := models.DatabaseConfig{
		DatabaseUser:     "cassandra",
		DatabasePassword: "cassandra",
		DatabaseHost:     "localhost:9042",
		DatabasePort:     9042,
		DatabaseKeyspace: "teste",
	}

	scylla := NewScylla(&config)

	t.Run("Expect to return success on scan map", func(t *testing.T) {
		mockScylla := mocks.NewIScylla(t)
		ctx := context.Background()

		mockScylla.On("ScanMap",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(
			nil,
		)

		stmt := `SELECT * FROM account WHERE email = ?`
		account := &models.Account{}
		results := map[string]interface{}{
			"email":        &account.Email,
			"full_number":  &account.FullNumber,
			"alias":        &account.Alias,
			"city":         &account.City,
			"date_time":    &account.DateTime,
			"district":     &account.District,
			"name":         &account.Name,
			"public_place": &account.PublicPlace,
			"status":       &account.Status,
			"zip_code":     &account.ZipCode,
		}

		err := scylla.ScanMap(ctx, stmt, results, "")

		assert.Nil(t, err)
	})

}
