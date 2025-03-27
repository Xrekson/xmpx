package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/xrekson/auction/pkg/model"
)

func Connect() *pgx.Conn {
	singleton, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	var greeting time.Time
	singleton.QueryRow(context.Background(), "select current_timestamp").Scan(&greeting)
	fmt.Println(greeting)
	return singleton
}

func GetAllusers() ([]model.User, []error) {
	db := Connect()
	var outUsers []model.User
	var outErrors []error

	if db == nil {
		log.Println("Database connection is nil")
		outErrors = append(outErrors, fmt.Errorf("database connection failed"))
		return nil, outErrors
	}

	rows, err := db.Query(context.Background(), "SELECT id, user_name, name, desx, about, password, email, phno, dob, type, createdat, updatedat FROM auction")
	if err != nil {
		outErrors = append(outErrors, err)
		return nil, outErrors
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.ID, &user.UserName, &user.Name, &user.Desx, &user.About,
			&user.Password, &user.Email, &user.Phno, &user.Dob, &user.Type,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			outErrors = append(outErrors, err)
			continue
		}
		outUsers = append(outUsers, user)
	}

	if rows.Err() != nil {
		outErrors = append(outErrors, rows.Err())
	}

	return outUsers, outErrors
}

func GetUser(user model.User) (*model.User, error) {
	db := Connect()

	if db == nil {
		log.Println("Database connection is nil")
		return nil, fmt.Errorf("database connection failed")
	}
	var userDB model.User

	err := db.QueryRow(context.Background(), "SELECT id, user_name, name, desx, about, password, email, phno, dob, type, createdat, updatedat FROM auction WHERE user_name ILIKE $1", user.UserName).Scan(&userDB.ID, &userDB.UserName, &userDB.Name, &userDB.Desx, &userDB.About,
		&userDB.Password, &userDB.Email, &userDB.Phno, &userDB.Dob, &userDB.Type,
		&userDB.CreatedAt, &userDB.UpdatedAt)
	if err != nil {
		return nil, err
	} else {
		return &userDB, nil
	}
}
