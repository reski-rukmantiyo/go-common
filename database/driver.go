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
	db.Client, db.Err = db.getClient()
	return db
}

func (db *DB) getClient() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta connect_timeout=5",
		db.host, db.username, db.password, db.name, db.port)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return client, err
}

func (db *DB) Error() bool {
	if db.Client != nil {
		return false
	}
	if db.Err != nil || db.Client == nil {
		db.Client, db.Err = db.getClient()
		return true
	}
	return false
}
