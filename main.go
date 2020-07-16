package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
)

func DownloadVideo(c *gin.Context) {
    url := c.PostForm("url")
    resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
    body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("found: %s %q\n", url, body)

    c.JSON(
        http.StatusOK,
        gin.H{"message": body},
    )
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
    v1 := router.Group("/api/v1/video")
    {
        v1.GET("/:name", Test)
        v1.POST("/:url", DownloadVideo)
    }
    router.Run(":8080")
}
