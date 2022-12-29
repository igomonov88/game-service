package codechallenge

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"

	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/internal/clients/codechallenge/contract"
	"githib.com/igomonov88/game-service/internal/middlewaries"
)

const getRandomNumberPath = "random"

var ErrInvalidConfig = errors.New("invalid config")

type Client struct {
	logger *zap.SugaredLogger
	host   string
	client *http.Client
}

func Must(client *Client, err error) *Client {
	if err != nil {
		panic(err)
	}
	return client
}

func NewClient(logger *zap.SugaredLogger, cfg Config) (*Client, error) {
	if len(cfg.Host) == 0 && cfg.RequestTimeout == 0 {
		return nil, ErrInvalidConfig
	}
	return &Client{
		logger: logger,
		host:   cfg.Host,
		client: &http.Client{Timeout: cfg.RequestTimeout},
	}, nil
}

func (c *Client) GenerateRandomNumber(ctx context.Context) int {
	reqID := middlewaries.RequestIDFromContext(ctx)
	responseData, err := c.doRequest(http.MethodGet, getRandomNumberPath)
	if err != nil {
		c.logger.Errorf("request_id: %v, failed to make request, got error: %v", reqID, err)
		return 0
	}

	var response contract.GetRandomNumberResponse
	if err = json.Unmarshal(responseData, &response); err != nil {
		c.logger.Errorf("request_id: %v, failed to unmarshal response, got error: %v", reqID, err)
		return 0
	}

	return response.RandomNumber
}

func (c *Client) doRequest(method string, methodPath string) ([]byte, error) {
	reqURL := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(methodPath),
	}

	req, err := http.NewRequest(method, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
