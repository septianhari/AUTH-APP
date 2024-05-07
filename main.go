package main

import (
	// "encoding/base64"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("avan-barno-tamvan")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var dbUsername = "sabarno"
var dbPassword = "opensesame"

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
			return
		}

		// Login with admin
		user.Role = "admin"

		// Create Token with Base64
		// data := []byte(user.Username + ":" + user.Password)
		// token := base64.StdEncoding.EncodeToString(data)

		// Create Token with JWT
		expiredTime := time.Now().Add(1 * time.Minute)
		payload := jwt.MapClaims{
			"role": user.Role,
			"exp":  expiredTime.Unix(),
			"iat":  time.Now().Unix(),
			"iss":  "auth-api",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": tokenString})
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
		// Parse Token:
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Validasi Token
		if !tkn.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Basic Authentication
		// username, _, ok := c.Request.BasicAuth()
		// if !ok {
		// 	c.JSON(401, gin.H{"error": "Unauthorized"})
		// 	c.Abort()
		// 	return
		// }

		// Send context
		c.Set("role", claims["role"])

		c.Next()
	})

	router.GET("/admin", func(c *gin.Context) {
		role := c.MustGet("role").(string)
		// Authorization
		if role != "admin" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.JSON(200, gin.H{"welcome ": role})
	})

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"username": c.MustGet("username")})
	})

	router.Run()

	// var inputUsername, inputPassword string

	// fmt.Print("Username: ")
	// fmt.Scan(&inputUsername)
	// fmt.Print("Password: ")
	// fmt.Scan(&inputPassword)

	// // Authentication:
	// if inputUsername == dbUsername && inputPassword == dbPassword {
	// 	fmt.Println("Login Success")
	// 	fmt.Println("Ini JWT(JSON Web Token): (dsadadaeqwewq213) username = " + inputUsername)
	// } else {
	// 	fmt.Println("Login Gagal")
	// }

	// // Authorization
	// if inputUsername == "admin" {
	// 	fmt.Println("Kamu bisa akses semua fitur")
	// } else {
	// 	fmt.Println("Kamu hanya bisa akses halaman user!")
	// }
}
