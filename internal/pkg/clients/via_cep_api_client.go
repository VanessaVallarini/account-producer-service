package clients

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sony/gobreaker"
)

const endUrl = "/json/"

type IViaCepApiClient interface {
	CallViaCepApi(ctx context.Context, ae models.ViaCepRequest) (*models.ViaCepResponse, error)
}

type ViaCepApiClient struct {
	cb  *gobreaker.CircuitBreaker
	url string
}

func NewViaCepApiClient(cfg *models.ViaCepConfig) *ViaCepApiClient {
	var st gobreaker.Settings
	st.Name = cfg.Name
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= uint32(cfg.MaxRetriesHttpRequest) && failureRatio >= cfg.MaxFailureRatio
	}

	return &ViaCepApiClient{
		cb:  gobreaker.NewCircuitBreaker(st),
		url: cfg.Url,
	}
}

func (v *ViaCepApiClient) CallViaCepApi(ctx context.Context, ae models.ViaCepRequest) (*models.ViaCepResponse, error) {

	body, err := v.cb.Execute(func() (interface{}, error) {
		url := v.url + ae.Cep + endUrl
		resp, err := http.Get(url)
		if err != nil {
			utils.Logger.Error("via cep api failed during call api: %v", err)
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	var viaCepResponse models.ViaCepResponse
	if err := json.Unmarshal(body.([]byte), &viaCepResponse); err != nil {
		utils.Logger.Error("via cep api failed during unmarshal return api: %v", err)
		return nil, err
	}

	return &viaCepResponse, nil
}
