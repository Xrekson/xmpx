package cmd

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/xrekson/auction/pkg/model"
  "github.com/go-pg/pg/v10/orm"
)

var dbConnection *pg.DB

func Connect() *pg.DB {
	opts := &pg.Options{
		Addr:     ":5432",
		User:     "app",
		Password: "Thundera@190",
	}
	if dbConnection ==nil {
     dbConnection = pg.Connect(opts)
	}
	return dbConnection
}

func GetAllusers() ([]model.User, error) {
	db := Connect()
	var outUsers []model.User

	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection failed")
	}

	rows := db.Model(&outUsers).Select()
	if rows != nil {
		return nil, fmt.Errorf("No user present in system!")
	}
	
	return outUsers, nil
}

func GetUser(user model.User) (*model.User, error) {
	db := Connect()

	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection failed")
	}
	var userDb model.User
		
	err := db.Model(userDb).Where("user.user_name = ?", user.UserName).Select()
	
	if err != nil {
		return nil, fmt.Errorf("User not found!")
	} else {
		return &userDb, nil
	}
}

func GetAllListings() ([]model.Listing, error){
	db := Connect()
	var outListing []model.Listing
	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection failed")
	}

	err := db.Model(&outListing).Select()
	fmt.Println("Hi",outListing, " ",err)
	if err != nil {
		return nil, fmt.Errorf("No listing present in system!")
	}
	
	return outListing, nil
}

func CreateSchema() error {
	db := Connect()
	models := []interface{}{
        (*model.Listing)(nil),
				(*model.User)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: true,
        })
				fmt.Println("create")
        if err != nil {
            return err
        }
    }
    return nil
}

