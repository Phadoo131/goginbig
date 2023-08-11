package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" 
	db "github.com/Phadoo131/goginbig/db/sqlc"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://bigsimpleapi:phadoo131@localhost:5444/?sslmode=disable"
)

var dbConn *sql.DB
var queries *db.Queries

func Init() {
	var err error
	dbConn, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		panic(err)
	}
	queries = db.New(dbConn)
}


func GetBook(c *gin.Context) {
	books, err := queries.ListinStore(c, db.ListinStoreParams{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		log.Fatal("query error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func AddBookInstore(c *gin.Context) {
	var newBook db.Instore

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := queries.CreateInstore(c, db.CreateInstoreParams{
		Book:      newBook.Book,
		Owner:     newBook.Owner,
		Bookcount: newBook.Bookcount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}


func CheckoutBook(c *gin.Context) {
	id := c.Query("id")

	instore, err := queries.GetInstoreForUpdate(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if instore.Bookcount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "There are no books available"})
		return
	}

	instore.Bookcount--
	updatedInstore, err := queries.UpdateinStore(c, db.UpdateinStoreParams{
		Book:      instore.Book,
		Owner:     instore.Owner,
		Bookcount: instore.Bookcount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedInstore)
}

func ReturnBook(c *gin.Context) {
	id := c.Query("id")

	instore, err := queries.GetInstoreForUpdate(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	instore.Bookcount++
	updatedInstore, err := queries.UpdateinStore(c, db.UpdateinStoreParams{
		Book:      instore.Book,
		Owner:     instore.Owner,
		Bookcount: instore.Bookcount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedInstore)
}
