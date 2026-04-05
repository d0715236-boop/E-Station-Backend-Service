package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Station представляет структуру данных зарядной станции
type Station struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	r := gin.Default()

	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to E-Station API",
			"website": "https://e-station.tech/",
		})
	})

	// Список станций (пример данных)
	r.GET("/api/stations", func(c *gin.Context) {
		stations := []Station{
			{ID: "1", Name: "Центральная АЗС", Status: "available"},
			{ID: "2", Name: "Парк Победы", Status: "busy"},
		}
		c.JSON(http.StatusOK, stations)
	})

	// Запуск сервера на порту 8080
	r.Run(":8080") 
}
