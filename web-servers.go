package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "Hello!")
	fmt.Fprint(w, "Hello!")
	fmt.Fprint(w, "Hello!")
	fmt.Fprint(w, "Hello!")
	fmt.Fprint(w, "Hello!")
	fmt.Fprint(w, "Hello!")
	fmt.Fprint(w, "Hello!")

	fmt.Println(r.UserAgent())
	fmt.Println(r.Host)
	fmt.Println(r.Header)
	fmt.Println(r.Cookies())

}
func main() {
	var h Hello

	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
