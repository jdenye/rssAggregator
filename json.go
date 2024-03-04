package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(sen http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil{
		log.Printf("Failed to marshal JSON response: %v", payload)
		sen.WriteHeader(500)
		return 
	}

	sen.Header().Add("Content-Type", "application/json")
	sen.WriteHeader(code)
	sen.Write(dat)
}