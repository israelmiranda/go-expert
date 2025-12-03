package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://viacep.com.br/ws/36025275/json/

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		res, err := http.Get(
			fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep),
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %v\n", err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %v\n", err)
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %v\n", err)
		}
		defer file.Close()

		_, err = fmt.Fprintf(file,
			"CEP: %s, Localidade: %s, UF: %s\n",
			data.Cep, data.Logradouro, data.Uf)
	}
}
