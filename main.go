package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"rpssl_service/config"
	"rpssl_service/routes"

)

func main() {
	n := negroni.Classic()
	n.UseHandler(routes.RegisterRoutes())

	port:= config.GetString("PORT","5000")

	log.Printf("starting server on port :%s",port)
	err := http.ListenAndServe(fmt.Sprintf(":%s",port), n)
	if err != nil {
		if r := recover(); r != nil {
			err = r.(error)
		}
		log.Fatal(err)
	}
}