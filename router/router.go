package router

import (
    "github.com/gin-gonic/gin"
    "Go-Gin-practice/apis"
    "Go-Gin-practice/initialize"
    "Go-Gin-practice/global"
    "Go-Gin-practice/core"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "Go-Gin-practice/docs" // docs is generated by Swag CLI, you have to import it.
)


func init_db() {
    initialize.DBTables()
    defer global.GVA_DB.Close()
    core.Init()
}

// @title Swagger Example API
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1/
func SetupRouter() *gin.Engine {
    init_db()
    router := gin.Default()
    api_v1 := router.Group("/api/v1")
    {
        shortener := api_v1.Group("/shortener")
        {
            shortener.GET("/:path", v1.GetURL)
            shortener.POST("", v1.CreateURL)
        }
    }
    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
    return router
}
