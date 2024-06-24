### Multithreaded CEP Search Challenge using APIs

In this repository, you will find a solution to the challenge of efficiently fetching address information based on postal codes (CEP), leveraging multithreading concepts and API integration.

#### Challenge Description

The goal is to simultaneously make requests to two distinct APIs that provide CEP information:

1. **BrasilAPI**: https://brasilapi.com.br/api/cep/v1/{cep}
2. **ViaCEP**: http://viacep.com.br/ws/{cep}/json/

The solution should:

- Perform parallel requests to both APIs.
- Compare response times and choose the faster response.
- Display the address data retrieved in the command line, along with the identification of the API that provided the faster response.
- Implement a 1-second timeout for each request. If one or both APIs do not respond within this time, a timeout error should be displayed.

#### Project Structure

```
.
├── cmd
│   └── main.go            # Entry point of the program
├── go.mod                 # Go module definition file
├── internals
│   └── cep
│       └── getcep.go      # Implementation of the logic to fetch CEPs
└── pkg
    ├── regex
    │   └── cep.go        # Functions for CEP validation
    └── request
        └── cep
            ├── brasilapi.go    # Client for BrasilAPI
            ├── viacep.go       # Client for ViaCEP
```

#### How to Use

1. **Prerequisites**:
   - Make sure you have Go installed on your machine.

2. **Installation**:
   - Clone this repository: `git clone https://github.com/guirialli/Go-Cep-Thread-Challenger`
   - Navigate to the project directory: `cd Go-Cep-Thread-Challenger`
   - Install project dependencies: `go mod tidy`

3. **Execution**:
   - Run the main program: `go run cmd/main.go`
   - Follow the instructions in the terminal to input the desired CEP according to the specified format.

#### Final Notes

This project demonstrates how to handle concurrency in Go to enhance the performance of API requests and manage timeouts to ensure application robustness. Feel free to explore, modify, and contribute improvements to this repository.
---
