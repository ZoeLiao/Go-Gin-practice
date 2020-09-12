package v1

import (
	"Go-Gin-practice/global"
	"Go-Gin-practice/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// Test example
// @Description test
// @name test name
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /shorteners [get]
func GetURLList(c *gin.Context) {
	var URLs models.URL
	res := global.GVA_DB.Find(&URLs)
	if res.Error != nil {
		global.GVA_LOG.Error("Error:", res.Error)
		c.JSON(
			http.StatusOK,
			gin.H{"res": "Fialed to get path list", "error": res.Error},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"path": res},
	)
}

// Test example
// @Description test
// @name test name
// @Accept  json
// @Produce  json
// @Param message body models.URL true "url"
// @Success 200 {string} string	"ok"
// @Router /shortener [post]
func CreateURL(c *gin.Context) {
	var URL models.URL
	c.BindJSON(&URL)
	url := URL.Url
	notHas := global.GVA_DB.Where("url = ?", url).Find(&URL).RecordNotFound()
	if !notHas {
		global.GVA_LOG.Info("url: ", url, " Already exists.")
		c.JSON(
			http.StatusOK,
			gin.H{"url": url, "res": "Already exists"},
		)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error(err)
		c.JSON(
			http.StatusOK,
			gin.H{"message": err},
		)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error(err)
		c.JSON(
			http.StatusOK,
			gin.H{"message": err},
		)
	} else {
		err = global.GVA_DB.Create(&URL).Error
		if err != nil {
			global.GVA_LOG.Error(err)
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
// @Param   path     path    string     true        "path"
// @Success 200 {string} string	"ok"
// @Router /shortener/{path} [get]
func GetURL(c *gin.Context) {
	path := c.Param("path")
	var URL models.URL
	res := global.GVA_DB.Where("path = ?", path).First(&URL)
	if res.Error != nil {
		global.GVA_LOG.Error("Error:", res.Error)
		c.JSON(
			http.StatusOK,
			gin.H{"path": path, "res": "Path not found", "error": res.Error},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"path": path, "url": URL.Url},
	)
}

// Test example
// @Description test
// @name test name
// @Accept  json
// @Produce  json
// @Param message body models.URL true "url"
// @Success 200 {string} string	"ok"
// @Router /shortener [patch]
func UpdateURL(c *gin.Context) {
	var URL models.URL
	c.BindJSON(&URL)
	path := URL.Path
	url := URL.Url
	notHas := global.GVA_DB.Where("path = ?", path).Find(&URL).RecordNotFound()
	if !notHas {
		global.GVA_LOG.Info("path: ", path, " path not found.")
		c.JSON(
			http.StatusOK,
			gin.H{"path": path, "res": "path not found"},
		)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error(err)
		c.JSON(
			http.StatusOK,
			gin.H{"res": err},
		)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error(err)
		c.JSON(
			http.StatusOK,
			gin.H{"res": err},
		)
	} else {
		err = global.GVA_DB.Update(&URL).Error
		if err != nil {
			global.GVA_LOG.Error(err)
			c.JSON(
				http.StatusOK,
				gin.H{"url": url, "res": "Failed to update URL"},
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
// @Param message body models.URL true "url"
// @Success 200 {string} string	"ok"
// @Router /shortener [delete]
func DeleteURL(c *gin.Context) {
	var URL models.URL
	c.BindJSON(&URL)
	path := URL.Path
	url := URL.Url
	notHas := global.GVA_DB.Where("path = ?", path).Find(&URL).RecordNotFound()
	if !notHas {
		global.GVA_LOG.Info("path: ", path, " path not found.")
		c.JSON(
			http.StatusOK,
			gin.H{"path": path, "res": "path not found"},
		)
		return
	}

	err := global.GVA_DB.Delete(&URL, 1).Error
	if err != nil {
		global.GVA_LOG.Error(err)
		c.JSON(
			http.StatusOK,
			gin.H{"url": url, "res": "Failed to delete URL"},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"url": url, "res": "Delete successfully"},
		)
	}
}
