package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB


type User struct {
	gorm.Model
	Name		string
	Email		string  `gorm:"type:varchar(100);unique_index"`
}
//TODO:: We should relate accounts to multiple users [Account have one to many relation with User]
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
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	//TODO:: Bring these configs to env file
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)
	new_db, err := gorm.Open(os.Getenv("db_type"), dbUri)

	db = new_db
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	//db.LogMode(true)
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
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T09:20:04-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Vancouver Ice Cream",
		NormalizedName:"Vancouver Ice Cream",
		Fee:2,
	}
	createTransaction(&transaction)

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
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T09:24:15-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Burnaby Taxi",
		NormalizedName:"Burnaby Taxi",
		Fee:0,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T11:41:10-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Purple Taxi",
		NormalizedName:"Purple Taxi",
		Fee:-7.5,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T12:57:06-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Station Cafe On Beatty",
		NormalizedName:"Station Cafe",
		Fee:-20.53,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T16:01:38-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Vancouver Gym Downtown",
		NormalizedName:"Vancouver Gym",
		Fee:-20.99,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T16:40:30-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"TaxSoft",
		NormalizedName:"TaxSoft",
		Fee:-44.8,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-01T16:40:30-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Phones R Us",
		NormalizedName:"Phones R Us",
		Fee:-41.29,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-02T09:50:29-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"Paypal *The Hobby",
		NormalizedName:"The Hobby",
		Fee:0,
	}
	createTransaction(&transaction)

	t, _ = time.Parse(time.RFC3339,"2019-04-02T11:43:11-07:00")
	transaction = Transaction{
		Account:account,
		Date:t,
		RawName:"DevTools",
		NormalizedName:"DevTools",
		Fee:-9.36,
	}
	createTransaction(&transaction)


	db.Close()
}

func createTransaction(transaction *Transaction) {
	db.Exec(`INSERT  INTO "transactions" ("created_at","updated_at","deleted_at","account_id","date","raw_name","normalized_name","fee") VALUES (NOW(),NOW(),NULL,
		`+ strconv.FormatInt(int64(transaction.Account.ID),10) + `1,NOW(),'`+ transaction.RawName + `','`+ transaction.NormalizedName + `',`+ strconv.FormatFloat(transaction.Fee,'f',6,64) + `) RETURNING "transactions"."id"
	`)
	db.Exec("UPDATE accounts SET balance = balance + (" + strconv.FormatFloat(transaction.Fee, 'f', 6, 64) +"), last_transaction = NOW() , updated_at= NOW() WHERE id=" + strconv.FormatInt(int64(transaction.Account.ID),10) + ";")
}


