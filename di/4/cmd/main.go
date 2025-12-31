package main

import (
	"context"
	"fmt"
	"time"
)

type address struct {
	API          string `json:"api"`
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func findAddressFromBrazilAPI(cep string, ch chan<- address) {
	client := BrasilApiClientProvider()
	res, err := client.FetchAddress(context.Background(), cep)
	if err != nil {
		return
	}

	ch <- address{
		API:          "brasilapi",
		Cep:          res.Cep,
		State:        res.State,
		City:         res.City,
		Neighborhood: res.Neighborhood,
		Street:       res.Street,
	}
}

func findAddressFromviaCEP(cep string, ch chan<- address) {
	client := ViaCepApiClientProvider()
	res, err := client.FetchAddress(context.Background(), cep)
	if err != nil {
		return
	}

	ch <- address{
		API:          "viacep",
		Cep:          res.Cep,
		State:        res.State,
		City:         res.City,
		Neighborhood: res.Neighborhood,
		Street:       res.Street,
	}
}

func main() {
	ch := make(chan address)
	cep := "01153000"

	go findAddressFromBrazilAPI(cep, ch)
	go findAddressFromviaCEP(cep, ch)

	select {
	case address := <-ch:
		fmt.Println("---------------------------------------")
		fmt.Printf("API: ** %s **\n", address.API)
		fmt.Printf("CEP: %s\n", address.Cep)
		fmt.Printf("State: %s\n", address.State)
		fmt.Printf("City: %s\n", address.City)
		fmt.Printf("Neighborhood: %s\n", address.Neighborhood)
		fmt.Printf("Street: %s\n", address.Street)
		fmt.Println("---------------------------------------")
	case <-time.After(1 * time.Second):
		fmt.Println("one second timeout exceeded!")
	}
}
