package initializers

import "api/models"


func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}