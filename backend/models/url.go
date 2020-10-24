package models

import (
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	Url  string `form:"url"`
	Path string
}
