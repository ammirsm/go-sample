package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", mainPage).Methods("GET")
	myRouter.HandleFunc("/cards", allCards).Methods("GET")
	myRouter.HandleFunc("/transactions/add_tag/", addTagForTransactions).Methods("POST")
	myRouter.HandleFunc("/transactions/delete_tag/{tag_id:[0-9]+}", deleteTagForTransactions).Methods("DELETE")
	myRouter.HandleFunc("/transactions/{account_id:[0-9]+}/", allTransactions).Methods("GET")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), myRouter))
}
