package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nikitakosatka/markdown-notes/pkg/repository"
	"log"
	"net/http"
	"time"
)

func Create(c *gin.Context) {
	var note repository.Note

	_ = c.BindJSON(&note)

	note.ID = uuid.New().String()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	if creationErr := repository.Create(&note); creationErr != nil {
		message := fmt.Sprintf("Error while creating new note into db. Reason: %v\n", creationErr)
		log.Printf(message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": message,
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

	if getErr := repository.GetAll(&notes); getErr != nil {
		message := fmt.Sprintf("Error while getting all notes, Reason: %v\n", getErr)
		log.Printf(message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": message,
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
	id := c.Param("id")
	note.ID = id

	if validationErr := c.ShouldBindJSON(&note); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Not enough args in request",
		})
		return
	}

	if updateErr := repository.Update(&note); updateErr != nil {
		message := fmt.Sprintf("Error. Reason: %v\n", updateErr)
		log.Printf(message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": message,
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

	if err := repository.Remove(note); err != nil {
		message := fmt.Sprintf("Error while deleting a single note, Reason: %v\n", err)
		log.Printf(message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Note deleted successfully",
	})
	return
}
