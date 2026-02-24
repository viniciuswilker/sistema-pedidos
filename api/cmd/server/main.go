package main

import (
	"fmt"

	"github.com/viniciuswilker/sistema-pedidos/configs"
)

func main() {
	fmt.Println("Rodando servidor na porta 8010")
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
