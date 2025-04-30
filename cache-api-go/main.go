package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const PORT = 8080

type CustomResponse struct {
	Value string `json:"value"`
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var value string

	// Check if cached
	if val, err := rdb.Get(ctx, id).Result(); err == nil {
		value = val
	} else {
		// Simulate slow DB
		value = fmt.Sprintf("data-for-id-%s", id)
		time.Sleep(5 * time.Second)

		// Save to Redis
		rdb.Set(ctx, id, value, 15*time.Second)
	}

	newData := CustomResponse{Value: value}
	data, err := json.Marshal(newData)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.Write(data)
}

func main() {
	initRedis()

	router := http.NewServeMux()
	router.HandleFunc("GET /data/{id}", getDataHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error at starting http server: %s\n", err)
	}

	fmt.Printf("Running server at port %d\n", PORT)
}
