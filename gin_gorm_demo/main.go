package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
)

type User struct {
	ID       uint `gorm:"primarykey;AUTO_INCREMENT"`
	Name     string
	Password string
}

func main() {
	// 数据库
	dsn := "root:root123@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "haha", Password: "123"})

	// 服务器
	server := gin.Default()
	server.Use(cors.Default())

	server.POST("/user/register", func(context *gin.Context) {
		u := User{}
		context.BindJSON(&u)
		// 查询用户名
		res := db.Where("name = ?", u.Name).First(&User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册失败，用户名已存在",
			})
		} else {
			db.Create(&u)
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册成功",
			})
		}
	})

	//登录

	server.POST("/user/login", func(context *gin.Context) {
		u := User{}
		context.BindJSON(&u)
		res := db.Where("name = ? AND password = ?", u.Name, u.Password).First(&User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "登录成功",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "登录失败,用户名或密码不正确",
			})
		}

	})
	server.Run(":8888")
}
