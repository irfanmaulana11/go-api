package models

import "time"

type User struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRules struct {
	Name     string `valid:"required~parameter is empty"`
	Email    string `valid:"required~parameter is empty,email~Please provide a valid email address"`
	Password string `valid:"required~parameter is empty,length(12|1000)~Password should be at least 12 characters long"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
