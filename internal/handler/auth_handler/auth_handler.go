package auth_handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Email string `json:"email" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
}

type HandlerHTTP struct{
	DB *sql.DB
}

func NewHandlerHTTP(db *sql.DB) *HandlerHTTP {
	return &HandlerHTTP{DB: db}
}

func(h *HandlerHTTP) Create(c *gin.Context) {
	var json Users
	
	fmt.Println("connect:==================", h.DB.Ping())

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("data request:===================", json)
		query := `INSERT INTO go_basic_db.users(email, pass) VALUES (?, ?)`

		result, err := h.DB.Exec(query, json.Email, json.Pass)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
			return
		}

		fmt.Println("data result:===================", result)

		insertedID, err := result.LastInsertId()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"messages": "inserted",
			"last_inserted_id": insertedID,
		})
		return

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}