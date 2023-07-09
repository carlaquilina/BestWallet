package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Address   string
	BirthDate time.Time
	Balance   float64
}
