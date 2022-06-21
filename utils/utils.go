package utils

import (
	"log"
	"net/http"
)

// ReturnJsonResponse function for returning movies data in JSON format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	logerr(res.Write(resMessage))
}

func logerr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}
