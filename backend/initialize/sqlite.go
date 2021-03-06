package initialize

import (
	"Go-Gin-practice/backend/global"
	"Go-Gin-practice/backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DBTables() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	} else {
		global.GVA_DB = db
	}

	// Migrate the schema
	db.AutoMigrate(models.URL{})
}
