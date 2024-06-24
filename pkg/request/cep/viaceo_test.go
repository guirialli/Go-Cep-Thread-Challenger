package cep_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/guirialli/desafio-mult-thread/pkg/regex"
	"github.com/guirialli/desafio-mult-thread/pkg/request/cep"
)

func TestGetViaCep(t *testing.T) {
	result, err := cep.GetViaCep("45570000")

	if err != nil {
		t.Error(err.Error())
	}

	if result.Cep != "45570-000" {
		err = fmt.Errorf("o cep buscadp nao corresponde ao esperado %s\n", result.Cep)
		t.Error(err.Error())
	}

	if result.Uf != "BA" {
		err = fmt.Errorf("O cep buscado não corresponde a uf da BA: %s\n", result.Uf)
		t.Error(err.Error())
	}

	if result.Localidade != "Ipiaú" {
		err := fmt.Errorf("O cep buscado não corresponde a cidade de Ipiaú: %s\n", result.Localidade)
		t.Error(err.Error())
	}

}

func TestToCep(t *testing.T) {
	result, err := cep.GetViaCep("45570000")
	if err != nil {
		t.Error(err.Error())
	}

	toCep := result.ToCep()

	if !regex.ValidCep(toCep.Cep) {
		err = fmt.Errorf("cep no formato invalido: %s\n", toCep.Cep)
		t.Error(err.Error())
	}

	if result.Bairro != toCep.Bairro || strings.Join(strings.Split(result.Cep, "-"), "") != toCep.Cep ||
		result.Localidade != toCep.Cidade || result.Uf != toCep.Estado ||
		result.Logradouro != toCep.Logradouro {
		err = fmt.Errorf("erro para converter:\nexpceted %v\nrecived: %v\n\n", result, toCep)
		t.Error(err.Error())
	}
}
