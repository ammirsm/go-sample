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
	LastTransaction	time.Time
	Balance			float64
}


// our initial migration function
func initialMigration() {
	db = openDb(db)
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{},&Account{},&Card{},&Transaction{}, &Tag{})
	initialSeedData()

}

func openDb(db *gorm.DB) (*gorm.DB){
	//TODO:: Bring these configs to env file
	new_db, err := gorm.Open("postgres", "host=localhost port=5432 user=amir dbname=wealth_ethical_w password=1234 sslmode=disable")
	db = new_db
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.LogMode(true)
	return db
}

func initialSeedData()  {
	db = openDb(db)


	var user = User{Name: "Amir", Email: "amir@fit-ro.com"}
	db.Create(&user)
	db.First(&user, user.ID)

	var account = Account{
		Number: "371851262565945",
		User: user,
		Balance:1371.51,
		LastTransaction:time.Now(),
	}
	db.Create(&account)
	db.First(&account, account.ID)

	var card = Card{
		Number:"4539011500139075",
		Account:account,
		IsActive:true,
		ExpirationDate:time.Date(2020, 2, 1, 12, 30, 0, 0, time.UTC),
	}
	db.Create(&card)
	db.First(&card,card.ID)

	t, _ := time.Parse(time.RFC3339,"2019-04-01T06:13:33-07:00")
	var transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"E-transfer Fee Reimbursement",
		NormalizedName:"E-transfer Fee",
		Fee:1.5,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T09:20:04-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Vancouver Ice Cream",
		NormalizedName:"Vancouver Ice Cream",
		Fee:0,
	}
	db.Create(&transaction)
	var tag Tag
	tag = Tag{Name:"Ice Cream"}
	db.Create(&tag)
	transaction.Tags = append(transaction.Tags,tag)
	db.Save(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T09:20:08-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Safeway",
		NormalizedName:"Safeway",
		Fee:0,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T09:24:15-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Burnaby Taxi",
		NormalizedName:"Burnaby Taxi",
		Fee:0,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T11:41:10-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Purple Taxi",
		NormalizedName:"Purple Taxi",
		Fee:-7.5,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T12:57:06-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Station Cafe On Beatty",
		NormalizedName:"Station Cafe",
		Fee:-20.53,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T16:01:38-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Vancouver Gym Downtown",
		NormalizedName:"Vancouver Gym",
		Fee:-20.99,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T16:40:30-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"TaxSoft",
		NormalizedName:"TaxSoft",
		Fee:-44.8,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T16:40:30-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Phones R Us",
		NormalizedName:"Phones R Us",
		Fee:-41.29,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-02T09:50:29-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Paypal *The Hobby",
		NormalizedName:"The Hobby",
		Fee:0,
	}
	db.Create(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-02T11:43:11-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"DevTools",
		NormalizedName:"DevTools",
		Fee:-9.36,
	}
	db.Create(&transaction)

	db.Close()

}