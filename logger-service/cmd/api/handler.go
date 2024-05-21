package main

import (
	"github.com/MuhamadAndre10/simple-microservices/logger-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (c *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	var requestPayload JSONPayload

	// read json into var
	_ = c.readJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err := c.Models.LogEntry.Insert(event)
	if err != nil {
		c.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	c.writeJSON(w, http.StatusOK, resp)
}
