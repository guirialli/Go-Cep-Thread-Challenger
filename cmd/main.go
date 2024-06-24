package main

import (
	"context"
	"fmt"
	"time"

	cepinternals "github.com/guirialli/desafio-mult-thread/internals/cep"
	"github.com/guirialli/desafio-mult-thread/pkg/regex"
	ceppkg "github.com/guirialli/desafio-mult-thread/pkg/request/cep"
)

func main() {
	var cepInput string

	fmt.Println("Por favor digite o cep a ser buscado: ")
	fmt.Scanf("%s", &cepInput)

	if !regex.ValidCep(cepInput) {
		err := fmt.Sprintf(`o cep "%s", não corresponde ao padrão esperado!`, cepInput)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cepChan := make(chan ceppkg.Cep)
	go cepinternals.GetBrasilCep(cepInput, cepChan)
	go cepinternals.GetViaCep(cepInput, cepChan)

	select {
	case result := <-cepChan:
		fmt.Println(result.String())
	case <-ctx.Done():
		fmt.Println("Timeout ao buscar cpf!")
	}

}
