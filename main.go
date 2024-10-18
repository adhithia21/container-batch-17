package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		sleepDuration, _ := strconv.Atoi(os.Getenv("SLEEP"))
		time.Sleep(time.Duration(sleepDuration) * time.Second)

		c.JSON(http.StatusOK, gin.H{
			"code":        200,
			"language":    "go version go1.20.6",
			"message":     "Server running on port 8000",
			"description": "ini aplikasi golang latihan",
		})
	})
	r.Run("0.0.0.0:8000")
}
