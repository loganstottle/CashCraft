package controller

import (
	"CashCraft/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" // I promise I will not leak an API key, I promise I will not leak an API key, I promise I will not leak an API key
)

func LoadEnv() { // Imports those hidden secret enviroment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env\n")
	}
}

/*
If you are self hosting you will need a .env file with
FINNHUB_API_KEY
DBUSER
DBPW
The database is MySql, and has been tested to work with MariaDB as well
If you are on the prod dev team, reach out to the ginger via email or discord
*/

func StartServer() {
	model.ConnectDatabase() // Model to connect to the database (We love MVC!)

	app := fiber.New() // Initialize the fiber app

<<<<<<< HEAD
<<<<<<< HEAD
	SetupHomeRoutes(app)           // Nice organization to keep the routes in seperate files (less merge errors)
	SetupAuthRoutes(app)           // Other halve of the routes
	SetupLeaderboardRoutes(app)
	app.Static("/*", "./view/404") // For if someone puts in a wrong link (We are not case sensitive though - ease of use)
=======
=======
>>>>>>> 66cabed62f2804abdd3f0f6a46e4244f50df7ac3
	SetupHomeRoutes(app) // Nice organization to keep the routes in seperate files (less merge errors)
	SetupAuthRoutes(app) // Other halve of the routes
	app.Static("/", "./view/home")
	app.Static("/register", "./view/register")
	app.Static("/login", "./view/login")
	// app.Static("/*", "./view/404") // For if someone puts in a wrong link (We are not case sensitive though - ease of use)
<<<<<<< HEAD
>>>>>>> 66cabed62f2804abdd3f0f6a46e4244f50df7ac3
=======
>>>>>>> 66cabed62f2804abdd3f0f6a46e4244f50df7ac3
	model.SetupStocks()
	app.Listen(":3000") // Starts the server to where people can connect to it
}
