package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  FirstName string
  LastName string
  Username string `gorm:"unique"`
  Phone string `gorm:"unique"`
  Email string `gorm:"unique"`
  Password string
}