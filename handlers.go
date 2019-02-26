package main

import (
	"encoding/json"
	"net/http"
)

var rests []Rest

func getRestsList(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(responseWriter).Encode(rests)
	if err != nil {
		consoleLog.Error("Error convert to JSON")
	}

}
