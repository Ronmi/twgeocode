package main

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Geocode100 struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}
type Geocode103 struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}
type Geocode106 struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}
type Geocode107 struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}
type Geocode112 struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}

type dbWriter struct {
	conn *gorm.DB
}

func (w *dbWriter) init() error {
	os.Remove("geocode.db")
	conn, err := gorm.Open(sqlite.Open("geocode.db"))
	if err != nil {
		return err
	}
	w.conn = conn

	err = conn.AutoMigrate(
		Geocode100{},
		Geocode103{},
		Geocode106{},
		Geocode107{},
		Geocode112{},
	)

	return err
}

func (w *dbWriter) save(data []*Info, typ string) error {
	table := "geocode" + typ
	log.Printf("%d recordes to be inserted to db table %s", len(data), table)
	err := w.conn.Table(table).Save(data).Error
	if err != nil {
		return err
	}
	log.Printf("inserted %d recordes to db table %s", len(data), table)

	return nil
}
