package Lesson21

import (
	"fmt"
	"log"
	"os"
)

type Handler interface {
	Handle(request string) string
}

type BaseHandler struct{}

func (h BaseHandler) Handle(request string) string {
	return "Обработан запрос: " + request
}

type LoggingDecorator struct {
	wrapped Handler
	logger  *log.Logger
}

func (d LoggingDecorator) Handle(request string) string {
	d.logger.Printf("Получен запрос: %s", request)
	response := d.wrapped.Handle(request)
	d.logger.Printf("Отправлен ответ: %s", response)
	return response
}

func main() {
	logger := log.New(os.Stdout, "[LOG] ", log.LstdFlags)

	base := BaseHandler{}

	logged := LoggingDecorator{
		wrapped: base,
		logger:  logger,
	}

	result := logged.Handle("GET /weather")
	fmt.Println(result)
}
