package main

import (
	"main/pkg/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const ()

func main() {
	bdb, err := sqlx.Connect("mysql", "***:***@tcp(127.0.0.1)/billing")
	if err != nil {
		panic(err)
	}

	defer bdb.Close()
	db, err := sqlx.Connect("mysql", "***:***@tcp(127.0.0.1)/crm")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	h := handlers.Database{BillDB: bdb, AppDB: db}
	app := fiber.New()

	handlers.SetRoutes(app, &h)
	app.Listen(":3000")

}
