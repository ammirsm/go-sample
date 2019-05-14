package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type JsonResponse struct {
	Code	int
	Description	string
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

func addTagForTransactions(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var tag Tag
	err := decoder.Decode(&tag)
	if err != nil {
		panic(err)
	}

	db = openDb(db)
	defer db.Close()

	var transaction Transaction
	db.First(&transaction, 1)
	var relatedTags []Tag
	db.Model(&transaction).Related(&relatedTags)

	for i := range relatedTags{
		if relatedTags[i].Name == tag.Name{
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(&JsonResponse{409,"You've added this tag to this transaction recently"})
			return
		}
	}

	//TODO: Should check the user that is requesting is the owner of transaction or not

	db.Create(&tag)

	json.NewEncoder(w).Encode(tag)
}

func deleteTagForTransactions(w http.ResponseWriter, r *http.Request) {
	db = openDb(db)
	defer db.Close()

	vars := mux.Vars(r)
	TagId := vars["tag_id"]
	TagIdInt, _ := strconv.ParseInt(TagId, 0, 64)


	var tag Tag
	db.First(&tag,TagIdInt)

	//TODO: Should check the user that is requesting is the owner of transaction or not

	db.Delete(&tag)

	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(&JsonResponse{410,"Deleted successfully"})
}

func mainPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world, this is wealth ethical API main page :)")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", mainPage).Methods("GET")
	myRouter.HandleFunc("/cards", allCards).Methods("GET")
	myRouter.HandleFunc("/transactions/add_tag/", addTagForTransactions).Methods("POST")
	myRouter.HandleFunc("/transactions/delete_tag/{tag_id:[0-9]+}", deleteTagForTransactions).Methods("DELETE")
	myRouter.HandleFunc("/transactions/{card_id:[0-9]+}/", allTransactions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
	fmt.Println("Server is up")
	initialMigration()
	handleRequests()
}