package initializers

import "github.com/mxnuchim/golang-auth-api/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}