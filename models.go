package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var db *gorm.DB


type User struct {
	gorm.Model
	Name		string
	Email		string  `gorm:"type:varchar(100);unique_index"`
}

type Account struct {
	gorm.Model
	Number		string
	User		User
	UserId		int	`gorm:"index;not null"`
}



// our initial migration function
func initialMigration() {
	db = openDb(db)
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{},&Account{},&Card{},&Transaction{}, &Tag{})
	//initialSeedData()

}

func openDb(db *gorm.DB) (*gorm.DB){
	new_db, err := gorm.Open("postgres", "host=localhost port=5432 user=amir dbname=wealth_ethical password=1234 sslmode=disable")
	db = new_db
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	return db
}

func initialSeedData()  {
	db = openDb(db)
	defer db.Close()

	db.Create(&User{Name: "Charles", Email: "charles2@rails.town"})
	var user User
	db.First(&user, 1)
	db.Create(&Account{Number: "371851262565945", User: user})
	var account Account
	db.First(&account, 1)
	var card Card
	db.Create(&Card{
		Number:"4539011500139075",
		Account:account,
		Balance:0,
		IsActive:true,
		ExpirationDate:time.Date(2020, 2, 1, 12, 30, 0, 0, time.UTC),
		LastTransaction:time.Now(),
	})
	db.First(&card,1)
	var transaction Transaction
	db.Create(&Transaction{
		Card:card,
		Date:time.Now(),
		RawName:"E-transfer Fee Reimbursement",
		NormalizedName:"Fee Reimbursement",
		Fee:12,
	})
	db.First(&transaction,1)
	fmt.Println(transaction)

	db.First(&transaction,1)
	var tag Tag
	tag = Tag{Name:"Coffee"}
	db.Create(&tag)
	transaction.Tags = append(transaction.Tags,tag)
	db.Save(&transaction)

}