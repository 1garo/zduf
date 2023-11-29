package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/1garo/zduf/database"
	"github.com/1garo/zduf/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

// Class schema.
type Class struct {
	ID        uint `json:"-"`
	Name string    `json:"name"      binding:"required"`
	Time      time.Time `json:"time"      binding:"required"`
	//TeacherId uuid.UUID `json:"teacherId"`
}

// Student schema.
type Student struct {
	ID   uint `json:"-"`
	FirstName string    `json:"first_name"      binding:"required"`
	LastName string    `json:"last_name"      binding:"required"`
	// Age validate if Student could be part of the a class based on his Age
	Age     uint    `json:"age"  binding:"required"`
	//Classes []Class `json:"-"`
	tableName string
}

func (s *Student) create() (int, error) {
	var id int
	err := conn.QueryRow(
		context.Background(),
		fmt.Sprintf("insert into %s(first_name, last_name, age) values ($1, $2, $3) returning id", s.tableName),
		s.FirstName, s.LastName, s.Age).Scan(&id)

	return id, err
}

func main() {
	r := gin.Default()
	urlExample := "postgres://zduf:admin@localhost:5432/zduf_db"
	conn, err := database.NewConnection(urlExample)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	http.SetRoutes(r)
	http.Configure()


	//r.POST("/class", func(c *gin.Context) {
	//	var class Class

	//	if err := c.BindJSON(&class); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}

	//	c.JSON(http.StatusOK, gin.H{
	//		"data": class,
	//	})
	//})

	//r.POST("/student", func(c *gin.Context) {
	//	var student Student

	//	student.tableName = "student"

	//	if err := c.BindJSON(&student); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}

	//	id, err := student.create()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}

	//	c.JSON(http.StatusOK, gin.H{
	//		"data": id,
	//	})
	//})

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
