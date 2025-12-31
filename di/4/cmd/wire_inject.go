//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/israelmiranda/go-expert/di/4/httpclients/brasilapi"
	"github.com/israelmiranda/go-expert/di/4/httpclients/viacep"
)

func newRestyClient() *resty.Client {
	return resty.New()
}

func BrasilApiClientProvider() *brasilapi.Client {
	wire.Build(newRestyClient, brasilapi.NewClient)
	return &brasilapi.Client{}
}

func newHttpClient() *http.Client {
	return &http.Client{}
}

func ViaCepApiClientProvider() *viacep.Client {
	wire.Build(newHttpClient, viacep.NewClient)
	return &viacep.Client{}
}
