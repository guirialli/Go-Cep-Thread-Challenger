package cep

import (
	ceppkg "github.com/guirialli/desafio-mult-thread/pkg/request/cep"
)

func GetViaCep(cep string, result chan<- ceppkg.Cep) {
	cepBuscado, err := ceppkg.GetViaCep(cep)
	if err != nil {
		panic(err)
	}
	result <- cepBuscado.ToCep()
}

func GetBrasilCep(cep string, result chan<- ceppkg.Cep) {
	cepBuscado, err := ceppkg.GetBrasilApiCep(cep)
	if err != nil {
		panic(err)
	}
	result <- cepBuscado.ToCep()
}
