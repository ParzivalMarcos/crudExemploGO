package main

import (
	"net/http"

	"crudExemploGO/routes"
)

func main() {
	routes.Rotas()
	http.ListenAndServe(":8000", nil)
}
