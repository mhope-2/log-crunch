package main

import (
	"fmt"
	"github.com/mhope-2/log-query/database/cassandra"
	handler "github.com/mhope-2/log-query/handlers"
	"log"
	"net/http"
)

func main() {

	h := handler.New(cassandra.GetLogsKeySpace())

	//http.HandleFunc("/query/", h.QueryMessage)
	http.HandleFunc("/message/", h.RetrieveMessage)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
