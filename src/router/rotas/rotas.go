package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) {
	for _, rota := range rotasUsuarios {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

}
