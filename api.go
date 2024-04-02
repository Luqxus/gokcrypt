package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type APIFunc func(writer http.ResponseWriter, request *http.Request) error

type APIServer struct{}

func NewAPIServer() *APIServer {
	return &APIServer{}
}

func (api *APIServer) Run() {}

func handler(fn APIFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(request.Context(), 30*time.Second)
		defer cancel()

		err := fn(writer, request.WithContext(ctx))
		if err != nil {
			log.Panic(err)
		}
	}
}
