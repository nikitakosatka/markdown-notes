package repository

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"
	"time"
)

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

type Note struct {
	ID        string    `json:"id" pg:",pk"`
	Title     string    `json:"title" binding:"required"`
	Body      string    `json:"body" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

func CreateNoteTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Note{}, opts)
	if createError != nil {
		log.Printf("Error while creating note table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Note table created")
	return nil
}
