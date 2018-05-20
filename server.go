package main

import (
	"go_test_sample/sample_gorilla/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	controller := &controllers.IndexController{}
	r.HandleFunc("/", controller.Get).Methods("GET")
	r.HandleFunc("/", controller.Post).Methods("POST")
	r.HandleFunc("/{task_id:[0-9]+}", controller.Put).Methods("PUT")

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":9000")
}
