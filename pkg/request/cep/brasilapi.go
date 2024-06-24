package cep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (b BrasilApiCep) ToCep() Cep {
	return Cep{
		Cep:        b.Cep,
		Logradouro: b.Street,
		Bairro:     b.Neighborhood,
		Cidade:     b.City,
		Estado:     b.State,
	}
}

func GetBrasilApiCep(cep string) (*BrasilApiCep, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	var cepBrasil BrasilApiCep

	req, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	json.Unmarshal(body, &cepBrasil)

	return &cepBrasil, err
}
