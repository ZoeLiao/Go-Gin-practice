package models

import (
    "github.com/jinzhu/gorm"
)


type URL struct {
    gorm.Model
    URL string `json:"url"`
}
