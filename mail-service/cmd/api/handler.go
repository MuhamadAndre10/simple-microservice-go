package main

import "net/http"

func (c *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type MailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload MailMessage

	err := c.readJSON(w, r, &requestPayload)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = c.Mailer.SendSMTPMessage(msg)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Message sent successfully to" + requestPayload.To,
	}

	c.writeJSON(w, http.StatusOK, payload)

}
