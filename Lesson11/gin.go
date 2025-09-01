package Lesson11

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name") // Получаем query-параметр "name"
		message := "Привет, " + name + "!"

		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.Run(":8080") // Запускаем сервер на порту 8080
}
