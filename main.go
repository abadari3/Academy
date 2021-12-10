package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.tmpl.html", nil)
	})

	router.GET("/3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index3.tmpl.html", nil)
	})

	router.GET("/4", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index4.tmpl.html", nil)
	})

	router.GET("/5", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index5.tmpl.html", nil)
	})

	router.Run(":" + port)
}
