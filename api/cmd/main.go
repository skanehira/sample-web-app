package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/skanehira/sample-web-app/api/app/server"
	"github.com/skanehira/sample-web-app/api/app/server/handler"
	"github.com/skanehira/sample-web-app/api/domain/service"
	"github.com/skanehira/sample-web-app/api/infra"
)

// ErrorResponse HTTPエラーのレスポンス構造体
type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := infra.NewUser(db)
	userService := service.NewUser(userRepo)
	h := handler.NewHandler(userService)
	server := server.New(e).RegistHandler(h)

	log.Fatal(server.Start())
}
