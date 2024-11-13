package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Dog struct {
	gorm.Model

	Name  string
	Breed string
}

func AddDog(c *fiber.Ctx) error {
	dog := Dog{
		Name:  c.FormValue("name"),
		Breed: c.FormValue("breed"),
	}

	if err := db.Create(&dog).Error; err != nil {
		return err
	}

	return DogRow(dog).Render(c.Context(), c)
}

func main() {
	newDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = newDB

	// Migrate the schema
	db.AutoMigrate(&Dog{})

	app := fiber.New()

	app.Static("/", "./public")

	app.Post("/dog", AddDog)
	app.Get("/table-rows", func(c *fiber.Ctx) error {
		dogs := make([]Dog, 0)

		if err := db.Find(&dogs).Error; err != nil {
			return err
		}

		return DogRows(dogs).Render(c.Context(), c)
	})

	app.Listen(":3000")
}
