package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.readJSON(w, r, &requestPayload)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate email
	user, err := c.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		c.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		c.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// use log service
	err = c.LogRequest("authentication", fmt.Sprintf("User %s authenticated", user.Email))
	if err != nil {
		c.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Welcome %s!", user.Email),
		Data:    user,
	}

	c.writeJSON(w, http.StatusOK, payload)
}

func (c *Config) LogRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
