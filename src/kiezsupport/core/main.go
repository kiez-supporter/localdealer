package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	"kiezsupport/core/api/models"
	"kiezsupport/core/bootstrap"
	"kiezsupport/core/db/gorm"
	"kiezsupport/core/router"
)

type Status struct {
	Status string `json:"status" xml:"status"`
}

func main() {
	fmt.Println("Welcome to the server")
	log.Printf("ENV : %s", bootstrap.App.ENV)

	e := router.New()

	// init database
	gorm.Init()
	//autoDropTables()
	//autoCreateTables()
	//autoMigrateTables()

	e.GET("/appHealth", AppHealth)

	_ = e.Start(":8000")
}

func AppHealth(c echo.Context) error {
	s := &Status{
		Status: "OK",
	}
	return c.JSON(http.StatusOK, s)
}

// autoCreateTables: create database tables using GORM
func autoCreateTables() {
	if !gorm.DBManager().HasTable(&models.Location{}) {
		gorm.DBManager().CreateTable(&models.Location{})
	}
}

// autoMigrateTables: migrate table columns using GORM
func autoMigrateTables() {
	gorm.DBManager().AutoMigrate(&models.Location{})
}

// auto drop tables on dev mode
func autoDropTables() {
	if bootstrap.App.ENV == "dev" {
		gorm.DBManager().DropTableIfExists(&models.Location{}, &models.Location{})
	}
}
