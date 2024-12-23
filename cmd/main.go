package main

import (
	"fmt"
	"sum/internal/handler/auth_handler"
	"sum/pkg/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// init db
	db := database.Init()

	fmt.Println("test===================", db.Ping())

	handler := auth_handler.NewHandlerHTTP(db)
	
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router := r.Group("/api")
	{
		router.POST("/register", handler.Create)
	}

	r.Run(":8080")
}
