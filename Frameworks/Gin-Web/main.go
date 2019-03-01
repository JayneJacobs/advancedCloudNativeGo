package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	engine.GET("/api/songs", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllSongs())
	})

	engine.POST("/api/songs", func(c *gin.Context) {
		var song Song
		if c.BindJSON(&song) == nil {
			id, created := CreateSong(song)
			if created {
				c.Header("location", "/api/songs/"+id)
				c.Status(http.StatusCreated)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	engine.GET("/api/songs/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		song, found := GetSong(id)

		if found {
			c.JSON(http.StatusOK, song)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	engine.PUT("/api/songs/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var song Song
		if c.BindJSON(&song) == nil {
			exists := UpdateSong(id, song)
			if exists {

				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusNotFound)
			}
		}
	})

	engine.DELETE("/api/songs/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		DeleteSong(id)
		c.Status(http.StatusOK)
	})
	// run server on Port
	// configuration for static files and templates
	engine.LoadHTMLGlob("./templates/*.html")
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	engine.GET("/..", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Advanced Cloud Native Go with Gin Framework",
		})
	})

	engine.Run(port())

}

// Define Port Environment Variable
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "6060"
	}
	return ":" + port
}
