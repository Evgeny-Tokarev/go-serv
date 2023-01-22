package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Form struct {
	Name    string
	Address string
}
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
func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//func formHandler(w http.ResponseWriter, r *http.Request) {
//	enableCors(&w)
//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//	w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, PATCH, DELETE, GET")
//	var f Form
//	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
//		fmt.Fprintf(w, "ParseForm() err: %v", err)
//		return
//	}
//	fmt.Fprintf(w, "POST request successful\n")
//	json1 := make(map[string]string)
//	json1["name"] = f.Name + " smth"
//	json1["address"] = f.Address + " smth"
//	jsonStr, err := json.Marshal(json1)
//	if err != nil {
//
//		log.Fatal(err)
//	}
//	w.Write(jsonStr)
//}
//
//func enableCors(w *http.ResponseWriter) {
//	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//}
//
//func InitServer() {
//	http.HandleFunc("/album", formHandler)
//	fmt.Printf("Starting backend at port 9000...\n")
//	go func() {
//		log.Fatal(http.ListenAndServe(":9000", nil))
	}()
}

//package main
//
//import (
//"github.com/gin-gonic/gin"
//_ "github.com/lib/pq"
//"net/http"
//)
//
//type Album struct {
//	ID     string  `json:"id"`
//	Title  string  `json:"title"`
//	Artist string  `json:"artist"`
//	Price  float64 `json:"price"`
//}
//
//var albums = []Album{
//	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}
//
//func getAlbums(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, albums)
//}
//func getAlbumByID(c *gin.Context) {
//	id := c.Param("id")
//
//	for _, a := range albums {
//		if a.ID == id {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
//}
//func postAlbums(c *gin.Context) {
//	var newAlbum Album
//
//	if err := c.BindJSON(&newAlbum); err != nil {
//		return
//	}
//
//	albums = append(albums, newAlbum)
//	c.IndentedJSON(http.StatusCreated, newAlbum)
//}
//
//func main() {
//	router := gin.Default()
//	router.GET("/albums", getAlbums)
//	router.GET("/albums/:id", getAlbumByID)
//	router.POST("/albums", postAlbums)
//	router.Run("localhost:8080")
//}
