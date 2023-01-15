package clients

import (
	"account-producer-service/internal/models"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViaCepApiClient(t *testing.T) {
	configViaCep := models.ViaCepConfig{
		Url:                   "https://viacep.com.br/ws/",
		MaxRetriesHttpRequest: 3,
		MaxFailureRatio:       0.6,
		Name:                  "HTTP GET",
	}

	viaCepApiClient := NewViaCepApiClient(&configViaCep)

	assert.NotNil(t, viaCepApiClient)
}

func TestCallViaCepApiReturnError(t *testing.T) {
	t.Run("Expect to return error during call via cep api and url is missing", func(t *testing.T) {
		configViaCep := models.ViaCepConfig{
			Url:                   "",
			MaxRetriesHttpRequest: 3,
			MaxFailureRatio:       0.6,
			Name:                  "HTTP GET",
		}
		viaCepApiClient := NewViaCepApiClient(&configViaCep)
		ctx := context.Background()

		request := models.ViaCepRequest{
			Cep: "01001-000",
		}

		response, err := viaCepApiClient.CallViaCepApi(ctx, request)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestCallViaCepApiReturnSuccess(t *testing.T) {
	t.Run("Expect to return success during send msg to create account", func(t *testing.T) {
		configViaCep := models.ViaCepConfig{
			Url:                   "https://viacep.com.br/ws/",
			MaxRetriesHttpRequest: 3,
			MaxFailureRatio:       0.6,
			Name:                  "HTTP GET",
		}
		viaCepApiClient := NewViaCepApiClient(&configViaCep)
		ctx := context.Background()

		request := models.ViaCepRequest{
			Cep: "01001-000",
		}

		response, err := viaCepApiClient.CallViaCepApi(ctx, request)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
