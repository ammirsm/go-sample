package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)


type Card struct {
	gorm.Model
	Number			string
	AccountId		int `gorm:"index;not null"`
	Account			Account
	Balance			float64
	IsActive		bool
	ExpirationDate	time.Time
	LastTransaction	time.Time
}


func allCards(w http.ResponseWriter, r *http.Request) {
	db = openDb(db)
	defer db.Close()
	var allCards []Card

	//TODO: we should got user from token but I just mock this up
	var user User
	db.First(&user,User{Email:"amir.saiedmehr@gmail.com"})

	var cards []Card
	db.Preload("Account.User").Preload("Tag").Find(&cards)


	//TODO: Should change this to some new query that get the cards from database
	for i := range cards{
		if cards[i].Account.UserId == int(user.ID){
			allCards = append(allCards,cards[i])
		}
	}

	json.NewEncoder(w).Encode(allCards)
}
