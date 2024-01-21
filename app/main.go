package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albuns", getAlbums)
	router.POST("/album", postAlbums)
	router.Run("localhost:8080")
}
