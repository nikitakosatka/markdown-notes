package config

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/nikitakosatka/markdown-notes/pkg/repository"
	"log"
	"os"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func Connect(dbConfig *DBConfig) *pg.DB {
	opts := &pg.Options{
		User:     dbConfig.Username,
		Password: dbConfig.Password,
		Addr:     fmt.Sprintf("%s:%s", dbConfig.Host, dbConfig.Port),
		Database: dbConfig.Name,
	}

	var db = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	err := repository.CreateNoteTable(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	repository.InitiateDB(db)

	return db
}
