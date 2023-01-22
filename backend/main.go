package main

import (
	"fmt"
	"github.com/Evgeny-Tokarev/go-serv/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type Album struct {
	//ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func getAlbumByTitle(c *gin.Context) {
	title := c.Param("title")
	if album, err := db.GetAlbumByTitle(title); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}
func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	fmt.Println(albums)
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/albums", getAlbums)
	router.GET("/albums/:title", getAlbumByTitle)
	router.POST("/albums", postAlbums)
	router.Run("localhost:9000")
}
