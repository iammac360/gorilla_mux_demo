package main

import (
	//"log"

	"github.com/codegangsta/negroni"
	//"net/http"
)

func main() {
	router := NewRouter()

	//log.Fatal(http.ListenAndServe(":8080", router))
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
