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
	AccountId		int	`gorm:"index;not null"`
	Account		Account
	Date		time.Time
	RawName		string
	NormalizedName	string
	Fee		float64
	Tags	[]Tag	`gorm:"foreignkey:TransactionId"`
}


func allTransactions(w http.ResponseWriter, r *http.Request) {

	//Pagination of transactions
	fromInt64, _ := strconv.ParseInt(r.URL.Query().Get("from"), 0, 64)
	limitInt64, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 0, 64)
	if int(limitInt64) == 0 || limitInt64 > paginationLimit {
		limitInt64 = paginationLimit
	}


	//Date picker
	fromDayAgo, _ := strconv.ParseInt(r.URL.Query().Get("from_day"), 0, 64)
	limitDays, _ := strconv.ParseInt(r.URL.Query().Get("limit_day"), 0, 64)
	if int(limitDays) == 0 || limitDays > dateLimit{
		limitDays = dateLimit
	}

	db = openDb(db)
	defer db.Close()

	vars := mux.Vars(r)
	AccountId := vars["account_id"]
	AccountIdInt, _ := strconv.ParseInt(AccountId, 0, 64)

	var account Account
	db.First(&account,int(AccountIdInt))

	var allTransactions []Transaction
	var transactions []Transaction

	fromDate := time.Now().Add(-24 * time.Duration(fromDayAgo) * time.Hour)
	toDate := fromDate.Add(-24 * time.Duration(limitDays) * time.Hour)
	db.Limit(int(limitInt64)).Offset(int(fromInt64)).Preload("Account").Preload("Tags").Where("date BETWEEN ? AND ?", toDate, fromDate).Find(&transactions)

	//TODO: Should change this to some new query that get the cards from database
	for i := range transactions{
		if transactions[i].AccountId == int(account.ID){
			allTransactions = append(allTransactions,transactions[i])
		}
	}

	json.NewEncoder(w).Encode(allTransactions)
}


func (transaction *Transaction) AfterSave(tx *gorm.DB) (err error) {
	var account Account
	tx.First(&account,transaction.AccountId)
	tx.Model(&account).Update("balance", account.Balance + transaction.Fee)
	tx.Model(&account).Update("last_transaction", transaction.Date)
	return

}