package main

import (
	s "algorithm/solution"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := s.CreateText(); err != nil {
		log.Print(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	res := s.Solution()
	fmt.Fprintf(w, "result= %v\n", res)
}
