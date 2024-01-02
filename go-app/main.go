package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Mulai aplikasi go")
	r := gin.Default()
	r.GET("/", Index)
	r.Run(":8000")
}

func Index(c *gin.Context) {
	_, err := ConnectDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Connected to database"})
}

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
