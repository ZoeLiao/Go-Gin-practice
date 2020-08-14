package v1

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
    "Go-Gin-practice/models"
    "Go-Gin-practice/global"
)


// Test example
// @Description test
// @name test name 
// @Accept  json
// @Produce  json
// @Param message body main.URL true "url"
// @Success 200 {string} string	"ok"
// @Router /shortner [post]
func shortener(c *gin.Context) {
    var URL models.URL
    c.BindJSON(&URL)
    url := URL.Url

    notHas := global.GVA_DB.Where("url = ?", url).Find(&URL).RecordNotFound()
    if !notHas {
        c.JSON(
            http.StatusOK,
            gin.H{"url": url, "res": "Already exists"},
        )
        return
    }

    resp, err := http.Get(url)
    if err != nil {
        fmt.Errorf("Error: %s", err)
        c.JSON(
            http.StatusOK,
            gin.H{"message": err},
        )
        return
    }

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
        err = global.GVA_DB.Create(&URL).Error
        if err != nil {
            c.JSON(
                http.StatusOK,
                gin.H{"url": url, "res": "Failed to create URL"},
            )
        } else {
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
// @Router /shortener/{path} [get]
func Test(c *gin.Context) {
    path := c.Param("path")
    var URL models.URL
    c.BindJSON(&URL)
    obj := global.GVA_DB.Where("path = ?", path).Find(&URL).RecordNotFound()
    if obj == nil {
        c.JSON(
            http.StatusOK,
            gin.H{"path": path, "res": "Path not found"},
        )
        return
    }
    c.JSON(
        http.StatusOK,
        gin.H{"path": path, "url": obj.url},
    )
}
