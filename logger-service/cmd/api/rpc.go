package main

import (
	"context"
	"github.com/MuhamadAndre10/simple-microservices/logger-service/data"
	"time"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, rsp *string) error {

	collection := client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	*rsp = "Processed payload via RPC: " + payload.Name

	return nil
}
