package main

import "net/http"

func handlerReadiness(sen http.ResponseWriter, res *http.Request){
	respondWithJSON(sen, 200, struct{}{})
}