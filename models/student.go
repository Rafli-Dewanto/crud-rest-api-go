package models

import (
	_ "gorm.io/gorm"
)

type Student struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Age         uint64 `json:"age"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
