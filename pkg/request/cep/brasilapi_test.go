package cep_test

import (
	"fmt"
	"testing"

	"github.com/guirialli/desafio-mult-thread/pkg/regex"
	"github.com/guirialli/desafio-mult-thread/pkg/request/cep"
)

func TestGetBrasilCep(t *testing.T) {
	result, err := cep.GetBrasilApiCep("45570000")
	if err != nil {
		t.Error(err.Error())
	}

	if result.Cep != "45570000" {
		err = fmt.Errorf("o cep buscadp nao corresponde ao esperado %s\n", result.Cep)
		t.Error(err.Error())
	}

	if result.State != "BA" {
		err = fmt.Errorf("O cep buscado não corresponde a uf da BA: %s\n", result.State)
		t.Error(err.Error())
	}

	if result.City != "Ipiaú" {
		err := fmt.Errorf("O cep buscado não corresponde a cidade de Ipiaú: %s\n", result.City)
		t.Error(err.Error())
	}
}

func TestToCepBrasilCep(t *testing.T) {
	result, err := cep.GetBrasilApiCep("45570000")
	if err != nil {
		t.Error(err.Error())
	}

	toCep := result.ToCep()

	if !regex.ValidCep(toCep.Cep) {
		err = fmt.Errorf("cep no formato invalido: %s\n", toCep.Cep)
		t.Error(err.Error())
	}

	if result.Neighborhood != toCep.Bairro || result.Cep != toCep.Cep ||
		result.City != toCep.Cidade || result.State != toCep.Estado ||
		result.Street != toCep.Logradouro {
		err = fmt.Errorf("erro para converter:\nexpceted %v\nrecived: %v\n\n", result, toCep)
		t.Error(err.Error())
	}
}
