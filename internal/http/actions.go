package http

import (
	"net/http"

	"github.com/1garo/zduf/database"
	"github.com/1garo/zduf/internal"
	"github.com/1garo/zduf/internal/teacher"
	"github.com/gin-gonic/gin"
)

var s *teacher.Service

func Configure() {
	s = &teacher.Service{
		Repository: teacher.Repository{
			Conn: database.Conn,
		},
	}
}

func CreateTeacher(c *gin.Context) {
	var teacher internal.Teacher

	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.Create(&teacher); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id": teacher.ID,
		},
	})
}
