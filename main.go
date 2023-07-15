package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/proullon/ramsql/driver"
	"log"
	"net/http"
	"strconv"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func GetBook(c *gin.Context) {
	books2 := []book{}
	i := c.Param("id")

	db := c.MustGet("db").(*sql.DB)

	if i == "" {
		rows, err := db.Query(`SELECT id, title, author, quantity FROM booksshelf`)
		if err != nil {
			log.Fatal("query error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var m book
			if err := rows.Scan(&m.ID, &m.Title, &m.Author, &m.Quantity); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "scan: " + err.Error()})
				return
			}
			books2 = append(books2, m)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, books2)
		return
	}

	id, err := strconv.Atoi(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := db.Query(`SELECT id, title, author, quantity FROM booksshelf where id =?`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var m book
		if err := rows.Scan(&m.ID, &m.Title, &m.Author, &m.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books2 = append(books2, m)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books2)
}

func CreateBook(c *gin.Context) {
	var newBook book
	// var newBooks []book
	// //Check body
	// c.Header("Content-Type", "application/json; charset=utf-8")
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(body))

	// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err := c.BindJSON(&newBook); err != nil {
		// c.JSON(http.StatusBadRequest, err.Error())
		// if err := c.BindJSON(&newBooks); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	// 	db := c.MustGet("db").(*sql.DB)

	// 	stmt, err := db.Prepare(`
	// 		INSERT INTO booksshelf(id, title, author, quantity)
	// 		VALUES (?, ?, ?, ?);
	// 	`)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	defer stmt.Close()

	// 	for _, b := range newBooks {
	// 		_, err := stmt.Exec(b.ID, b.Title, b.Author, b.Quantity)
	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 	}

	// 	c.JSON(http.StatusCreated, newBooks)
	// 	return
	// }

	db := c.MustGet("db").(*sql.DB)

	stmt, err := db.Prepare(`
		INSERT INTO booksshelf(id, title, author, quantity)
		VALUES (?, ?, ?, ?);
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}


func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*sql.DB)

	row := db.QueryRow(`SELECT id, title, author, quantity
	FROM booksshelf WHERE id=?`, id)
	m := book{}
	err := row.Scan(&m.ID, &m.Title, &m.Author, &m.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)
}

func CheckoutBook(c *gin.Context) {
	id := c.Query("id")
	db := c.MustGet("db").(*sql.DB)

	row := db.QueryRow(`SELECT id, title, author, quantity FROM booksshelf WHERE id=?`, id)
	m := book{}
	err := row.Scan(&m.ID, &m.Title, &m.Author, &m.Quantity)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if m.Quantity <= 0{
		c.JSON(http.StatusBadRequest, gin.H{"message":"There are no book availble"})
		return
	}

	m.Quantity -= 1

	_, err = db.Exec(`UPDATE booksshelf SET quantity = ? WHERE id = ?`, m.Quantity, m.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)

}

func ReturnBook(c *gin.Context) {
	id := c.Query("id")
	db := c.MustGet("db").(*sql.DB)

	row := db.QueryRow(`SELECT id, title, author, quantity FROM booksshelf WHERE id=?`, id)
	m := book{}
	err := row.Scan(&m.ID, &m.Title, &m.Author, &m.Quantity)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	m.Quantity += 1

	_, err = db.Exec(`UPDATE booksshelf SET quantity = ? WHERE id = ?`, m.Quantity, m.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)

}

var db *sql.DB

func conn() {
	var err error
	db, err = sql.Open("ramsql", "booksshelf")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn()

	createTb := `
	CREATE TABLE IF NOT EXISTS booksshelf (
	id INT AUTO_INCREMENT,
	title TEXT NOT NULL UNIQUE,
	author TEXT NOT NULL,
	quantity INT NOT NULL,
	PRIMARY KEY (id)
	);
	`

	_, err := db.Exec(createTb)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/books", GetBook)
	router.POST("/books", CreateBook)
	router.GET("/books/:id", GetBookByID)
	router.PATCH("/checkout", CheckoutBook)
	router.PATCH("/return", ReturnBook)
	router.Run("localhost:2566")
}