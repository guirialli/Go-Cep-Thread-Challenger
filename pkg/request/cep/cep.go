package cep

import "fmt"

type Cep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"localidade"`
	Estado     string `json:"uf"`
}

func (c *Cep) String() string {
	return fmt.Sprintf("{Cep: %s, Logradouro: %s, Bairro: %s, Cidade: %s, Estado: %s}\n",
		c.Cep, c.Logradouro, c.Bairro, c.Cidade, c.Estado)
}

type CepInterface interface {
	ToCep() Cep
}
