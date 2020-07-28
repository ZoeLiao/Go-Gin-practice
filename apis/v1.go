package v1

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
    "Go-Gin-practice/models"
)


// Test example
// @Description test
// @name test name 
// @Accept  json
// @Produce  json
// @Param message body main.URL true "url"
// @Success 200 {string} string	"ok"
// @Router /download [post]
func DownloadURL(c *gin.Context) {
    var URL models.URL
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
                gin.H{"url": url, "body": body},
            )
        }
    }
}

// Test example
// @Description test
// @name test name 
// @Accept  json
// @Produce  json
// @Param   name     path    string     true        "name"
// @Success 200 {string} string	"ok"
// @Router /download/{name} [get]
func Test(c *gin.Context) {
    name := c.Param("name")
    c.JSON(
        http.StatusOK,
        gin.H{"message": name},
    )
}
