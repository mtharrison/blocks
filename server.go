package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/solve", solveHandle)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func solveHandle(w http.ResponseWriter, r *http.Request) {

	puzzle := Puzzle{}
	d := json.NewDecoder(r.Body)
	d.Decode(&puzzle)

	solution := solve(puzzle)

	e := json.NewEncoder(w)
	e.Encode(solution)
}
