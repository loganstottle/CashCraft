package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LoginInput struct { // This defines a struct for username and password, to keep them bound together
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupAuthRoutes(app *fiber.App) { // This connects the different subdomains of the sites to functions that distribute the views
	app.Get("/register", GetRegister)         //         Grabs the registration website
	app.Get("/login", GetLogin)               //               Grabs the login website
	app.Post("/register", RegisterHandler)    //    Allows user to request a new account
	app.Post("/login", LoginHandler)          //          Lets the user attempt to login
	app.Post("/logout", LogoutHandler)        //        Lets the user logout of their session
	app.Get("/me", AuthMiddleware, MeHandler) // *TODO*
}

/* The MVC we are using essentially means that we have multiple views that the user sees
   We have models which are similar to objects in java. We need to easily add multiple stocks and multiple users (why they are models)
   Then we have controllers, which are used to connect everything together - this system enables the project to be split throughout
   multiple files - which makes us able to all work on the project at the same time!
*/

func GetRegister(c *fiber.Ctx) error {
	var user model.User
	if err := model.DB.First(&user, "username = ? and session_token = ?", c.Cookies("username"), c.Cookies("session_token")).Error; err == nil {
		return c.Redirect("/")
	} // This code takes someone who is already logged in and takes them to the home page

	return c.Render("./view/register/index.html", fiber.Map{}) // This is what returns the view to the user
}

func GetLogin(c *fiber.Ctx) error {
	var user model.User
	if err := model.DB.First(&user, "username = ? and session_token = ?", c.Cookies("username"), c.Cookies("session_token")).Error; err == nil {
		return c.Redirect("/")
	} // This code takes someone who is already logged in and takes them to the home page

	return c.Render("./view/login/index.html", fiber.Map{})
}

func RegisterHandler(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil { // Checks that the user input does not have an error when parsed
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"}) // If error, tell them the input was incorrect
	}

	var user model.User

	// Prevents multiple accounts with the same username
	if err := model.DB.First(&user, "username = ?", input.Username).Error; err == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Account already exists"})
	}

	user.Username = input.Username
	user.Password = model.HashPassword(input.Password) // We hash the password to keep data safe!

	user.Cash = 100000.00 // Sets the initial amount of cash to 100k
	/* We determined this value because we figured we have to have big enough numbers for people to get excited,
	while having a reasonable amount of money (No billionairs)
	*/

	sessionToken := uuid.New().String() // This is so you dont have to login every time you open the website
	user.SessionToken = sessionToken
	model.DB.Create(&user) // We add a user to the database with this

	c.Cookie(&fiber.Cookie{ // Pretty generic session token cookie
		Name:     "username",
		Value:    input.Username,
		HTTPOnly: true,
		Secure:   true,     // Encryption
		SameSite: "Strict", // No other sites can steal the cookie (we don't want walmart.com to steal your login info)
	}) // I actually found the bug where we originally didn't have this - feel a bit pround
	// # I <3 Bugtesting

	c.Cookie(&fiber.Cookie{ // Pretty generic username cookie
		Name:     "session_token",
		Value:    sessionToken,
		HTTPOnly: true,
		Secure:   true,     // Encryption
		SameSite: "Strict", // No other sites can steal the cookie (we don't want walmart.com to steal your login info)
	})

	return c.JSON(fiber.Map{"message": "Logged in!"}) // This means everything has worked!
}

func LoginHandler(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil { //  Again, just checks and makes sure user input is valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user model.User
	if err := model.DB.First(&user, "username = ?", input.Username).Error; err != nil { // Checks that a user exists with that username
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if user.Password != model.HashPassword(input.Password) { // Checks that the users hashed password matches with the hashed input password
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	sessionToken := uuid.New().String() // New login on a new device, we make a new session token
	user.SessionToken = sessionToken    // We want to kick off any old sessions of the app
	model.DB.Save(&user)                // This serves to keep the account safe

	c.Cookie(&fiber.Cookie{ // Pretty generic username cookie
		Name:     "username",
		Value:    input.Username,
		HTTPOnly: true,
		Secure:   true,     // Encryption
		SameSite: "Strict", // No other sites can steal the cookie (we don't want walmart.com to steal your login info)
	})

	c.Cookie(&fiber.Cookie{ // Pretty generic session token cookie
		Name:     "session_token",
		Value:    sessionToken,
		HTTPOnly: true,
		Secure:   true,     // Encryption
		SameSite: "Strict", // No other sites can steal the cookie (we don't want walmart.com to steal your login info)
	})

	return c.JSON(fiber.Map{"message": "Logged in!"}) // Everything has worked!
}

/*
This does make me a bit curious, why do we not when using register
register them and then log them in using the log in method
In reality I know it is because I didnt think through the process at the start
but that is a good idea for the future, and there is no reason to switch it
*/

func LogoutHandler(c *fiber.Ctx) error { // This is really smart, originally we just cleared the device cookies
	sessionToken := c.Cookies("session_token")
	model.DB.Model(&model.User{}).Where("session_token = ?", sessionToken).Update("session_token", nil)
	// By deleteing the session token in the db if someone grabbed the session token while you where connected
	// then you sign out, they couldnt access your account

	// We still do delete the cookies though
	c.ClearCookie("username")
	c.ClearCookie("session_token")
	return c.JSON(fiber.Map{"message": "Logged out"})
}

func MeHandler(c *fiber.Ctx) error { // Basic function to return info about yourself
	user := c.Locals("user").(model.User)
	return c.JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
	})
}
