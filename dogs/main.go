package main

import (
	"strconv"

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

func DeleteDog(c *fiber.Ctx) error {
	dog := Dog{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	dog.ID = uint(id)

	if err := db.Delete(&dog).Error; err != nil {
		return err
	}

	return nil
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
	app.Delete("/dog/:id", DeleteDog)

	app.Listen(":3000")
}
