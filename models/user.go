package models

import "gorm.io/gorm"

// User Model
type User struct {
	gorm.Model
	Email    string
	Password string
	LoanID   string
}
