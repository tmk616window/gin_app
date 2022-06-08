package main

import (
	"github.com/gin-gonic/gin"
)

type PrintJob struct {
	Format string `json:"format" binding:"required"`
	InvoiceId int `json:"invoiceId" binding:"required,gte=0"`
    JobId int `json:"jobId" binding:"gte=0"`
}

func main() {
	router := gin.Default()
	router.POST("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{"user": "test"})
	})
    router.Run(":3001")
}
