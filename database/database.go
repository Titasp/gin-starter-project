package database

import (
	"github.com/boltdb/bolt"
)

var db *bolt.DB

func Init() (err error) {
	db, err = bolt.Open("./data.db", 0600, nil)
	return err
}

func GetDB() *bolt.DB {
	return db
}
