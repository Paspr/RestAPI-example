package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title  string   `json:"title"`
	Genre  []string `json:"genre"`
	Actors []string `json:"actors"`
}

var movies []Movie

func init() {
	movies = make([]Movie, 0)
}

func NewMovieHandler(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	movies = append(movies, movie)
	c.JSON(http.StatusOK, movie)
}

func ListMoviesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func UpdateMovieHandler(c *gin.Context) {
	id := c.Param("title")
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	index := -1
	for i := 0; i < len(movies); i++ {
		if movies[i].Title == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Movie is not found"})
		return
	}
	movies[index] = movie
	c.JSON(http.StatusOK, movie)
}

func DeleteMovieHandler(c *gin.Context) {
	id := c.Param("title")
	index := -1
	for i := 0; i < len(movies); i++ {
		if movies[i].Title == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Movie is not found"})
		return
	}
	movies = append(movies[:index], movies[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Movie has been deleted"})
}

func main() {
	router := gin.Default()
	router.POST("/movies", NewMovieHandler)
	router.GET("/movies", ListMoviesHandler)
	router.PUT("/movies/:title", UpdateMovieHandler)
	router.DELETE("/movies/:title", DeleteMovieHandler)
	router.Run()
}
