package main

import (
	s "algorithm/solution"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if res := s.Solution(); res != nil {
		fmt.Fprintf(w, "result= %v\n", res)
	}
}
