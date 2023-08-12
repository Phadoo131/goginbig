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

	secured.POST("/createaccount", service.CreateUserAccount)
	secured.POST("/deleteaccount", service.DeleteUserAccount)
	secured.GET("/account/admin", service.GetUserAccount)
	secured.GET("/accounts/admin", service.GetAllUserAccounts)
	secured.GET("/books", service.GetBookInstore)
	secured.POST("/addbook", service.AddBookInstore)
	secured.PATCH("/checkout", service.CheckoutBook)
	secured.PATCH("/return", service.ReturnBook)
	router.Run("localhost:2566")
}


func AdaptHTTPHandlerToGin(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}