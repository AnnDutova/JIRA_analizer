package utils

import (
	"encoding/json"
	"net/http"
)

type Href struct {
	Href string `json:"href"`
}

type BaseMessage struct {
	status  string
	message string
	name    string
	links   map[string]Href
}

func Message(status bool, message string, name string, url string) map[string]interface{} {
	links := make(map[string]Href)
	links["self"] = Href{Href: "http://localhost:8000" + url}

	return map[string]interface{}{"status": status, "message": message, "name": name, "_links": links}
}

func RespondAny(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
