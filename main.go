package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, albums)
	})

	router.GET("/albums/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Loop over the list of albums, looking for
		// an album whose ID value matches the parameter.
		for _, a := range albums {
			if a.ID == id {
				ctx.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	})

	router.POST("/albums", func(ctx *gin.Context) {
		var newAlbum album

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := ctx.BindJSON(&newAlbum); err != nil {
			return
		}

		// Add the new album to the slice.
		albums = append(albums, newAlbum)
		ctx.IndentedJSON(http.StatusCreated, newAlbum)
	})

	router.Run(":8080")
}
