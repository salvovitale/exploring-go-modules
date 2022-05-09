package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// importing multiple versions of the same package
	"rsc.io/quote"
	quoteV2 "rsc.io/quote/v2"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	// using multiple versions of the same package
	fmt.Println(quote.Hello())
	fmt.Println(quoteV2.HelloV2())
	fmt.Println(quoteV3.HelloV3())

	rtr := mux.NewRouter()

	rtr.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("Topic: " + vars["topic"]))
	})

	http.Handle("/", rtr)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
