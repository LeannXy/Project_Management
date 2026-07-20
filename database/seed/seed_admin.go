package seed

import (
	"log"

	"github.com/LeannXy/Project_Management/config"
	"github.com/LeannXy/Project_Management/models"
	"github.com/LeannXy/Project_Management/utils"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name: "Super Admin",
		Email: "admin@example.com",
		Password: password,
		Role: "admin",
		PublicID: uuid.New(),
	}
	if err :=config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed too seed admin",err)
	} else{
		log.Println("Admin user seeded")
	}
}