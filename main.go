package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
)

type URL struct {
    URL string `json:"url" form:"url"`
}

func DownloadURL(c *gin.Context) {
    var URL URL
    c.BindJSON(&URL)
    url := URL.URL
    resp, err := http.Get(url)
    if err != nil {
        fmt.Errorf("Error: %s", err)
        c.JSON(
            http.StatusOK,
            gin.H{"message": err},
        )
    } else {
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Errorf("Error: %s", err)
            c.JSON(
                http.StatusOK,
                gin.H{"message": err},
            )
        } else {
            fmt.Printf("found: %s %q\n", url, body)
            c.JSON(
                http.StatusOK,
                gin.H{"message": body},
            )
        }
    }
}

func Test(c *gin.Context) {
    name := c.Param("name")
    c.JSON(
        http.StatusOK,
        gin.H{"message": name},
    )
}

func main() {
    router := gin.Default()
    v1 := router.Group("/api/v1/download")
    {
        v1.GET("/:name", Test)
        v1.POST("/", DownloadURL)
    }
    router.Run(":8080")
}
