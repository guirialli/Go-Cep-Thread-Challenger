package cep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Viacep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Ibge       string `json:"ibge"`
	Gia        string `json:"gia"`
	Ddd        string `json:"ddd"`
	Siafi      string `json:"siafi"`
}

func (c Viacep) ToCep() Cep {
	return Cep{
		Cep:        strings.Join(strings.Split(c.Cep, "-"), ""),
		Logradouro: c.Logradouro,
		Bairro:     c.Bairro,
		Cidade:     c.Localidade,
		Estado:     c.Uf,
	}
}
func GetViaCep(cep string) (*Viacep, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep)
	var cepResponse Viacep

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &cepResponse)
	return &cepResponse, nil
}
