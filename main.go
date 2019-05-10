package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func allTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CardId := vars["card_id"]
	fmt.Println(CardId)
	fmt.Fprintf(w, "All Transaction is here")
}

func allCards(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Cards is here")
}

func mainPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world, this is wealth ethical API main page :)")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", mainPage).Methods("GET")
	myRouter.HandleFunc("/cards", allCards).Methods("GET")
	myRouter.HandleFunc("/transactions/{card_id:[0-9]+}/", allTransactions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
	fmt.Println("Server is up")
	handleRequests()
}