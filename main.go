package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var id int = 1

func main() {
	r := gin.Default()

	// Create a new book
	r.POST("/books", func(c *gin.Context) {
		var newBook Book
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newBook.ID = id
		id++
		books = append(books, newBook)
		c.JSON(http.StatusCreated, newBook)
	})

	// Get all books
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is homepage")
	})

	// Get a single book by ID
	r.GET("/books/:id", func(c *gin.Context) {
		paramID := c.Param("id")
		for _, book := range books {
			if paramID == strconv.Itoa(book.ID) {
				c.JSON(http.StatusOK, book)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// Update a book by ID
	r.PUT("/books/:id", func(c *gin.Context) {
		paramID := c.Param("id")
		var updatedBook Book
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for index, book := range books {
			if paramID == strconv.Itoa(book.ID) {
				books[index] = updatedBook
				c.JSON(http.StatusOK, updatedBook)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// Delete a book by ID
	r.DELETE("/books/:id", func(c *gin.Context) {
		paramID := c.Param("id")
		for index, book := range books {
			if paramID == strconv.Itoa(book.ID) {
				books = append(books[:index], books[index+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	r.Run(":7777")
}
