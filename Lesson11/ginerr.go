package Lesson11

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DivideRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func main() {
	r := gin.Default()

	r.POST("/divide", func(c *gin.Context) {
		var req DivideRequest

		// Читаем JSON из тела запроса
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неправильный формат JSON"})
			return
		}

		// Проверка деления на ноль
		if req.B == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "деление на ноль"})
			return
		}

		// Возвращаем результат
		result := req.A / req.B
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.Run(":8080")
}
