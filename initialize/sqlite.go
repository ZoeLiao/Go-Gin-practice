package initialize

import (
    "github.com/jinzhu/gorm"
    "Go-Gin-practice/models"
    "Go-Gin-practice/global"
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
