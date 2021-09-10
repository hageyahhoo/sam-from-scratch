package main

import (
	"net/http"
)

func init() {
	Init()
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w)
	w.Write([]byte(GetAllServants()))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(getHandler))
	http.ListenAndServe(":8080", mux)
}
