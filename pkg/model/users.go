package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	UserName  string    `json:"userName" db:"user_name"`
	Name      string    `json:"name" db:"name"`
	Desx      string    `json:"desx" db:"desx"`
	About     string    `json:"about" db:"about"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Phno      string    `json:"phno" db:"phno"`
	Dob       time.Time `json:"dob" db:"dob"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
