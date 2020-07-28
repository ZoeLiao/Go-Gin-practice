package initialize

import (
    "github.com/jinzhu/gorm"
    "Go-Gin-practice/models"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DBTables() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(models.URL{})
}
