package main

import (
	"encoding/json"
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
	db = openDb(db)
	defer db.Close()
	var all_cards []Card

	// we should got user from token but I just mock this up
	var user User
	db.First(&user,User{Email:"amir.saiedmehr@gmail.com"})

	var cards []Card
	db.Preload("Account.User").Find(&cards)


	//TODO: Should change this to some new query that get the cards from database
	for i,_ := range cards{
		if cards[i].Account.UserId == int(user.ID){
			all_cards = append(all_cards,cards[i])
		}
	}

	json.NewEncoder(w).Encode(all_cards)
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
	initialMigration()
	handleRequests()
}