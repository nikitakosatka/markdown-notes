package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nikitakosatka/markdown-notes/pkg/repository"
	"log"
	"net/http"
	"time"
)

func Create(c *gin.Context) {
	var note repository.Note

	c.BindJSON(&note)

	note.ID = uuid.New().String()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	err := repository.Create(&note)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Note created Successfully",
	})
}

func GetAll(c *gin.Context) {
	var notes []repository.Note
	err := repository.GetAll(&notes)

	if err != nil {
		log.Printf("Error while getting all notes, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All notes",
		"data":    notes,
	})
}

func Read(c *gin.Context) {
	id := c.Param("id")

	note, err := repository.Read(id)
	if err != nil {
		log.Printf("Error while getting a single note, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single note",
		"data":    note,
	})
}

func Update(c *gin.Context) {
	var note repository.Note
	c.BindJSON(&note)

	err := repository.Update(&note)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Note Edited Successfully",
	})
}

func Remove(c *gin.Context) {
	id := c.Param("id")
	note := &repository.Note{ID: id}

	err := repository.Remove(note)
	if err != nil {
		log.Printf("Error while deleting a single note, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Note deleted successfully",
	})
	return
}
