package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)


type Tag struct {
	gorm.Model
	TransactionId	int	`gorm:"index;not null"`
	Name		string
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

