package main

import (
    "Go-Gin-practice/router"
)


func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
