package main

import (
	"database/sql"
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
	service "github.com/Phadoo131/goginbig/service"
	auth "github.com/Phadoo131/goginbig/service"
	middleware "github.com/Phadoo131/goginbig/service"
	db "github.com/Phadoo131/goginbig/db/sqlc"
	
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://bigsimpleapi:phadoo131@localhost:5444/?sslmode=disable"
)

var dbConn *sql.DB
var queries *db.Queries
var err error


func main() {
	dbConn, err = sql.Open(DBDriver, DBSource)
	if err != nil{
		log.Fatal("Cannot access database")
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", dbConn)
		c.Next()
	})

	router.POST("/login", gin.WrapH(http.HandlerFunc(auth.LoginHandler)))

	secured := router.Group("/", middleware.AuthorizationMiddleware)

	secured.GET("/books", service.GetBook)
	secured.POST("/addbook", service.AddBookInstore)
	secured.PATCH("/checkout", service.CheckoutBook)
	secured.PATCH("/return", service.ReturnBook)
	router.Run("localhost:2566")
}
