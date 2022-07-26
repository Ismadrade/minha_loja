package main

import (
	"net/http"

	"github.com/ismadrade/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
