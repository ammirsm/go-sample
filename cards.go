package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)


type Card struct {
	gorm.Model
	Number			string
	AccountId		int `gorm:"index;not null"`
	Account			Account
	IsActive		bool
	ExpirationDate	time.Time
}


func allCards(w http.ResponseWriter, r *http.Request) {
	fromInt64, _ := strconv.ParseInt(r.URL.Query().Get("from"), 0, 64)
	limitInt64, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 0, 64)
	if int(limitInt64) == 0 || limitInt64 > paginationLimit {
		limitInt64 = paginationLimit
	}

	db = openDb(db)
	defer db.Close()
	var allCards []Card

	//TODO: we should got user from token but I just mock this up
	var user User
	db.First(&user,User{Email:"amir@fit-ro" +
		".com"})

	var cards []Card
	db.Limit(int(limitInt64)).Offset(int(fromInt64)).Preload("Account.User").Find(&cards)

	//TODO: Should change this to some new query that get the cards from database
	for i := range cards{
		if cards[i].Account.UserId == int(user.ID){
			allCards = append(allCards,cards[i])
		}
	}

	json.NewEncoder(w).Encode(allCards)
}
