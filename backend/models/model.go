package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string `json:"Name"`
	Author string `json:"Author"`
	//Genres []string `json:"Genres"`
	//Url    string   `json:"Url"`
}
