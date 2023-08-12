package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "github.com/Phadoo131/goginbig/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

//Accounts_________________________________________________________

func CreateUserAccount(c *gin.Context){
	var newAccount db.Account

	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc, err := queries.CreateAccount(c, newAccount.Owner)
	if err != nil{
		log.Fatal("Create Account - Status Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, acc)
}

func DeleteUserAccount(c *gin.Context){
    id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = queries.DeleteAccount(c, idInt)
	if err != nil {
		log.Fatal("Delete Account - Status Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}


func GetUserAccount(c *gin.Context){
	id := c.Param("id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	acc, err := queries.GetAccount(c, idInt)
	if err != nil{
		log.Fatal("See Account - Status Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, acc)
}

func GetAllUserAccounts(c *gin.Context){
	acc, err := queries.ListAccounts(c)
	if err != nil{
		log.Fatal("See all accounts - Status Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, acc)
}

//Instore_________________________________________________________

func GetBookInstore(c *gin.Context) {
	books, err := queries.GetInstore(c)
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
	updatedInstore, err := queries.UpdateInstore(c, db.UpdateInstoreParams{
		Book:      instore.Book,
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
	updatedInstore, err := queries.UpdateInstore(c, db.UpdateInstoreParams{
		Book:      instore.Book,
		Bookcount: instore.Bookcount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedInstore)
}
