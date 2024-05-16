package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Broker service is running",
	}

	_ = c.writeJSON(w, http.StatusOK, payload)
}

func (c *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := c.readJSON(w, r, &requestPayload)
	if err != nil {
		c.errorJSON(w, err)
	}

	switch requestPayload.Action {
	case "auth":
		c.authenticate(w, requestPayload.Auth)
	default:
		c.errorJSON(w, errors.New("unknown action"))
	}
}

func (c *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we will send to the authentication service
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	req, err := http.NewRequest("POST", "http://auth-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.errorJSON(w, err)
	}

	defer resp.Body.Close()
	// make sure we get back the correct status code.

	if resp.StatusCode == http.StatusUnauthorized {
		c.errorJSON(w, errors.New(resp.Status), http.StatusUnauthorized)
		return
	} else if resp.StatusCode != http.StatusOK {
		c.errorJSON(w, errors.New("unexpected status code"), http.StatusInternalServerError)
		return
	}

	// create variable we will read resp body
	var jsonFromService jsonResponse

	// decode the response json
	err = json.NewDecoder(resp.Body).Decode(&jsonFromService)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		c.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	payload := new(jsonResponse)
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	c.writeJSON(w, http.StatusAccepted, payload)
}
