package main

import (
	"fmt"
	"github.com/mhope-2/log-query/database/cassandra"
	"github.com/mhope-2/log-query/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	keySpace, err := cassandra.GetLogsKeySpace()
	if err != nil {
		log.Fatalf("Error connecting to logs Keyspace: %v", err)
	}

	h := handler.New(keySpace)

	//http.HandleFunc("/query/", h.QueryMessage)
	http.HandleFunc("/messages/", h.RetrieveMessage)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
