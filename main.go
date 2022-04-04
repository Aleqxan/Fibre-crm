package main

import (
	"fmt"

	"github.com/Aleqxan/Fibre-crm/database"
	"github.com/Aleqxan/Fibre-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead:id", lead.GetLead)
	app.Post("/api/v1/lead:id", lead.NewLead)
	app.Delete("/api/v1/lead:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
