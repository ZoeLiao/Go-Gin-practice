package models

import (
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	Url  string `json:"url"`
	Path string `json:"path"`
}
