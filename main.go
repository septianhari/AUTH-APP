package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var dbUsername = "admin"
var dbPassword = "12345"

func main() {
	fmt.Println("Auth App!")

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Authentication:
		if user.Username != dbUsername || user.Password != dbPassword {
			c.JSON(401, gin.H{"error": "Invalid username or password"})
		}

		// Create JWT
		//jwt.
		c.JSON(200, gin.H{"token": "dsadadaeqwewq213"})
	})

	// Middleware
	router.Use(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token := authHeader[len("Bearer "):]
		if token != "dsadadaeqwewq213" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	})

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"username": "admin"})
	})

	router.Run()

	// var inputUsername, inputPassword string

	// fmt.Print("Username: ")
	// fmt.Scan(&inputUsername)
	// fmt.Print("Password: ")
	// fmt.Scan(&inputPassword)

	// // Authentication:
	// if inputUsername == dbUsername && inputPassword == dbPassword {
	// fmt.Println("Login Success")
	// fmt.Println("Ini JWT(JSON Web Token): (dsadadaeqwewq213) username = " + inputUsername)
	// } else {
	// fmt.Println("Login Gagal")
	// }

	// // Authorization
	// if inputUsername == "admin" {
	// fmt.Println("Kamu bisa akses semua fitur")
	// } else {
	// fmt.Println("Kamu hanya bisa akses halaman user!")
	// }
}
