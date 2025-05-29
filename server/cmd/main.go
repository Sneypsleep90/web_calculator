package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"web_calculator/internal/calculationService"
	"web_calculator/internal/db"
	"web_calculator/internal/handlers"
)

func main() {
	database, err := db.InitDB()

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	e := echo.New()

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculation)

	e.PATCH("/calculations/:id", calcHandlers.PatchCalculation)
	e.DELETE("calculations/:id", calcHandlers.DeleteCalculation)

	e.Start("localhost:8080")

}
