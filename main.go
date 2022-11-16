package main

import (
	"fmt"
	"github.com/fabasoad/go-hello-world/pkg/calculator"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("%d", calculator.Add(3, 5))))
		if err != nil {
			panic("Panic!")
		}
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error in HTTP", err)
	}
}
