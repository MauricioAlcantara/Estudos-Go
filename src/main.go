package main

import (
	"Estudos-Go/src/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
