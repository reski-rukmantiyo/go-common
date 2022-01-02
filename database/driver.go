package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	username string
	password string
	host     string
	port     string
	name     string
	Err      error
	Client   *gorm.DB
}

func NewDatabase(username, password, host, port, name string) *DB {
	db := &DB{
		username: username,
		password: password,
		host:     host,
		port:     port,
		name:     name,
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, username, password, name, port)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		db.Err = err
	}
	db.Client = client
	return db
}
