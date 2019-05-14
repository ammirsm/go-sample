package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)


type Transaction struct {
	gorm.Model
	CardId		int	`gorm:"index;not null"`
	Card		Card
	Date		time.Time
	RawName		string
	NormalizedName	string
	Fee		float64
	Tags	[]Tag	`gorm:"foreignkey:TransactionId"`
}


func allTransactions(w http.ResponseWriter, r *http.Request) {
	//TODO: Should handle date picker
	//TODO: Should handle pagination

	db = openDb(db)
	defer db.Close()

	vars := mux.Vars(r)
	CardId := vars["card_id"]
	CardIdInt, _ := strconv.ParseInt(CardId, 0, 64)

	var card Card
	db.First(&card,int(CardIdInt))

	var allTransactions []Transaction
	var transactions []Transaction

	db.Preload("Card").Preload("Tags").Find(&transactions)

	//TODO: Should change this to some new query that get the cards from database
	for i := range transactions{
		if transactions[i].CardId == int(card.ID){
			allTransactions = append(allTransactions,transactions[i])
		}
	}

	json.NewEncoder(w).Encode(allTransactions)
}