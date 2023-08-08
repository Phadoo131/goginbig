package main

import (
	"database/sql"
	"testing"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"bytes"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/proullon/ramsql/driver"
	"github.com/Phadoo131/goginbig/util"
	"github.com/stretchr/testify/require"
)

func createTestDB() *sql.DB {
	db, err := sql.Open("ramsql", "test_db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.POST("/books", CreateBook)
	router.GET("/books", GetBook)
	return router
}

var createTb = `
CREATE TABLE IF NOT EXISTS booksshelf (
	id INT AUTO_INCREMENT,
	title TEXT NOT NULL UNIQUE,
	author TEXT NOT NULL,
	quantity INT NOT NULL,
	PRIMARY KEY (id)
);
`

func TestCreateBook(t *testing.T) {
	db := createTestDB()
	defer db.Close()

	_, err := db.Exec(createTb)
	require.NoError(t, err)

	router := setupRouter(db)

	newBook := book{
		Title:    util.RandomTitle(),
		Author:   util.RandomAuthor(),
		Quantity: util.RandomQuantity(),
	}

	newBookJSON, err := json.Marshal(newBook)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(newBookJSON))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	var responseBook book
	err = json.Unmarshal(w.Body.Bytes(), &responseBook)
	require.NoError(t, err)

	require.Equal(t, newBook.Title, responseBook.Title)
	require.Equal(t, newBook.Author, responseBook.Author)
	require.Equal(t, newBook.Quantity, responseBook.Quantity)
}


func TestGetBook(t *testing.T) {
	db := createTestDB()
	defer db.Close()

	_, err := db.Exec(createTb)
	require.NoError(t, err)

	router := setupRouter(db)

	// Create a new book using the CreateBook handler.
	newBook := book{
		Title:    util.RandomTitle(),
		Author:   util.RandomAuthor(),
		Quantity: util.RandomQuantity(),
	}

	createBookJSON, err := json.Marshal(newBook)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(createBookJSON))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	// Fetch all books using the GetBook handler.
	req, err = http.NewRequest("GET", "/books", nil)
	require.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var responseBooks []book
	err = json.Unmarshal(w.Body.Bytes(), &responseBooks)
	require.NoError(t, err)

	// Ensure that the response contains the book we created.
	require.True(t, isBookInSlice(newBook, responseBooks))

	// Fetch the specific book by ID using the GetBook handler.
	req, err = http.NewRequest("GET", "/books/"+responseBooks[0].ID, nil)
	require.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var responseBook book
	err = json.Unmarshal(w.Body.Bytes(), &responseBook)
	require.NoError(t, err)

	// Ensure that the response matches the book we created earlier.
	require.Equal(t, newBook.Title, responseBook.Title)
	require.Equal(t, newBook.Author, responseBook.Author)
	require.Equal(t, newBook.Quantity, responseBook.Quantity)
}

func isBookInSlice(bookToFind book, books []book) bool {
	for _, b := range books {
		if b.ID == bookToFind.ID {
			return true
		}
	}
	return false
}