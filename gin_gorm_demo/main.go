package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

type User struct {
	ID       uint   `gorm:"primarykey;AUTO_INCREMENT"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	// Database
	dsn := "root:root123@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "haha", Password: "123"})

	// Server
	server := gin.Default()
	server.Use(cors.Default())

	server.Static("/", "./")

	server.POST("/user/register", func(context *gin.Context) {
		fmt.Println("Received registration request")
		u := User{}
		if err := context.BindJSON(&u); err != nil {
			log.Println("Request format error:", err)
			context.JSON(http.StatusBadRequest, gin.H{"msg": "Request format error"})
			return
		}
		// Check username
		var existingUser User
		res := db.Where("name = ?", u.Name).First(&existingUser)
		if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
			log.Println("Query failed:", res.Error)
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "Query failed"})
			return
		}
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{"msg": "Registration failed, username already exists"})
		} else {
			if err := db.Create(&u).Error; err != nil {
				log.Println("Registration failed:", err)
				context.JSON(http.StatusInternalServerError, gin.H{"msg": "Registration failed"})
				return
			}
			log.Println("Registration successful, user info:", u)
			context.JSON(http.StatusOK, gin.H{"msg": "Registration successful"})
		}
	})

	// Login
	server.POST("/user/login", func(context *gin.Context) {
		u := User{}
		if err := context.BindJSON(&u); err != nil {
			log.Println("Request format error:", err)
			context.JSON(http.StatusBadRequest, gin.H{"msg": "Request format error"})
			return
		}
		res := db.Where("name = ? AND password = ?", u.Name, u.Password).First(&User{})
		if res.Error != nil {
			log.Println("Query failed:", res.Error)
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "Query failed"})
			return
		}
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "Login successful",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "Login failed, incorrect username or password",
			})
		}
	})

	server.Run(":8888")
}
