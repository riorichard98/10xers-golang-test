package main

import (
	"10xers/domain"
	"10xers/handler"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	// Connect to SQLite database
    db, err := gorm.Open(sqlite.Open("10xers-phones"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // AutoMigrate the phone table
    db.AutoMigrate(&domain.Phone{})

	// create new PhoneDomain
	phoneDomain := domain.NewPhoneDomain(db)

	// creating new PhoneHandler
	phoneHandler := handler.NewPhoneHandler(phoneDomain)

	// Create a new Echo instance
	e := echo.New() 

	// Middleware for basic authentication 
	// hardcoded for now to provide validation for admin only who knows the username and password
    e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
        // Check if username and password are valid
        if username == "admin" && password == "123" {
            return true, nil
        }
        return false, nil
    }))

	phoneRoute := e.Group("/phones") // grouping route for phones route

	phoneRoute.GET("",phoneHandler.SearchPhone) // endpoint("/phones") for search handphones
	phoneRoute.POST("",phoneHandler.CreatePhone) // for creating new phone in db
	phoneRoute.PATCH("/:id",phoneHandler.UpdatePhone) // for updating phone by id
	phoneRoute.DELETE("/:id",phoneHandler.DeletePhoneByID) // for deleting phone by id
    
	// start the server in port 8080
    e.Logger.Fatal(e.Start(":8080"))
} 