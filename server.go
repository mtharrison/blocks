package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/solve", solveHandle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func solveHandle(w http.ResponseWriter, r *http.Request) {

	puzzle := Puzzle{}
	d := json.NewDecoder(r.Body)
	d.Decode(&puzzle)

	solution := solve(puzzle)

	e := json.NewEncoder(w)
	e.Encode(solution)
}
