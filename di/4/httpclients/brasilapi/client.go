package brasilapi

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func NewClient(client *resty.Client) *Client {
	client.SetHeader("Content-Type", "application/json")
	client.SetBaseURL("https://brasilapi.com.br")

	return &Client{client}
}

func (c *Client) FetchAddress(ctx context.Context, cep string) (*BrazilAPIResponse, error) {
	var response BrazilAPIResponse
	res, err := c.client.R().
		SetContext(ctx).
		SetResult(&response).
		Get(fmt.Sprintf("/api/cep/v1/%s", cep))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("code: %d, message: %s", res.StatusCode(), res.Body())
	}

	return &response, nil
}
