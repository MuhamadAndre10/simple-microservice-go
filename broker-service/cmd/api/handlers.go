package main

import (
	"net/http"
)

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Broker service is running",
	}

	_ = c.writeJSON(w, http.StatusOK, payload)
}
