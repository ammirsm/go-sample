package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
import "strconv"



func allTransactions(w http.ResponseWriter, r *http.Request) {
	db = openDb(db)
	defer db.Close()

	vars := mux.Vars(r)
	CardId := vars["card_id"]
	CardIdInt, _ := strconv.ParseInt(CardId, 0, 64)

	var card Card
	db.First(&card,int(CardIdInt))

	var allTransactions []Transaction
	var transactions []Transaction

	db.Preload("Card").Find(&transactions)

	//TODO: Should change this to some new query that get the cards from database
	for i := range transactions{
		if transactions[i].CardId == int(card.ID){
			allTransactions = append(allTransactions,transactions[i])
		}
	}

	json.NewEncoder(w).Encode(allTransactions)
}

func allCards(w http.ResponseWriter, r *http.Request) {
	db = openDb(db)
	defer db.Close()
	var allCards []Card

	// we should got user from token but I just mock this up
	var user User
	db.First(&user,User{Email:"amir.saiedmehr@gmail.com"})

	var cards []Card
	db.Preload("Account.User").Find(&cards)


	//TODO: Should change this to some new query that get the cards from database
	for i := range cards{
		if cards[i].Account.UserId == int(user.ID){
			allCards = append(allCards,cards[i])
		}
	}

	json.NewEncoder(w).Encode(allCards)
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