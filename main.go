package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()

	fmt.Println("Rodando a API na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
