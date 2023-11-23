package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Class schema.
type Class struct {
	ID        uuid.UUID `json:"-"`
	Name      string    `json:"name"      binding:"required"`
	Time      time.Time `json:"time"      binding:"required"`
	TeacherId uuid.UUID `json:"teacherId"`
}

// Student schema.
type Student struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name" binding:"required"`
	// Age validate if Student could be part of the a class based on his Age
	Age     uint    `json:"age"  binding:"required"`
	Classes []Class `json:"-"`
}

// Teacher schema.
type Teacher struct {
	ID      uuid.UUID `json:"-"`
	Name    string    `json:"name" binding:"required"`
	Age     uint      `json:"age"  binding:"required"`
	Teaches []Class   `json:"-"`
}

func main() {
	r := gin.Default()
	urlExample := "postgres://zduf:admin@localhost:5432/zduf_db"
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r.POST("/teacher", func(c *gin.Context) {
		var teacher Teacher

		if err := c.BindJSON(&teacher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": teacher,
		})
	})

	r.POST("/class", func(c *gin.Context) {
		var class Class

		if err := c.BindJSON(&class); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": class,
		})
	})

	r.POST("/student", func(c *gin.Context) {
		var student Student

		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": student,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
