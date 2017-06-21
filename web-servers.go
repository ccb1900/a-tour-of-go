package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServerHTTP(w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "Hello!")

}
func main() {
	//var h Hello

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
