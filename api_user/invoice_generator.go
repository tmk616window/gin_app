package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Invoice struct {
    InvoiceId int `json:"invoiceId"`
    CustomerId int `json:"customerId" binding:"required,gte=0"`
    Price int `json:"price" binding:"required,gte=0"`
    Description string `json:"description" binding:"required"`
}
type PrintJob struct {
    JobId int `json:"jobId"`
    InvoiceId int `json:"invoiceId"`
    Format string `json:"format"`
}

type AuthSuccess struct {
    Password int `json:"password"`
    UserName int `json:"username"`
}


func createPrintJob(invoiceId int) {
    client := resty.New()

    var p PrintJob
    // Call PrinterService via RESTful interface



    res, err := client.R().
        SetBody(PrintJob{Format: "A4", InvoiceId: invoiceId}).
        SetResult(&p).
        EnableTrace().
        Post("http://localhost:3001/print-jobs")

        fmt.Println("jdepwjopdewあああああああ")
        fmt.Println("  Status Code:", res.StatusCode())
        fmt.Println("  Status     :", res.Status())
        fmt.Println("jdepwjopdewあああああああ", err)
        fmt.Println("jdepwjopdewあああああああ")
        fmt.Println("jdepwjopdewあああああああ")
        fmt.Println("jdepwjopdewあああああああ")
        fmt.Println("jdepwjopdewあああああああ")

    resp, err := client.R().
    EnableTrace().
    Get("https://httpbin.org/get")

// Explore response object
    fmt.Println("Response Info:")
    fmt.Println("  Error      :", err)
    fmt.Println("  Status Code:", resp.StatusCode())
    fmt.Println("  Status     :", resp.Status())
    fmt.Println("  Proto      :", resp.Proto())
    fmt.Println("  Time       :", resp.Time())
    fmt.Println("  Received At:", resp.ReceivedAt())
    fmt.Println("  Body       :\n", resp)
    fmt.Println()


    if err != nil {
        log.Println("InvoiceGenerator: unable to connect PrinterService")
    }
    log.Printf("InvoiceGenerator: created print job #%v via PrinterService", p.JobId)
}
func main() {
    router := gin.Default()
    router.POST("/invoices", func(c *gin.Context) {
        var iv Invoice
        if err := c.ShouldBindJSON(&iv); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input!"})
            return
        }
        log.Println("InvoiceGenerator: creating new invoice...")
        rand.Seed(time.Now().UnixNano())
        iv.InvoiceId = rand.Intn(1000)
        log.Printf("InvoiceGenerator: created invoice #%v", iv.InvoiceId)
        
        createPrintJob(iv.InvoiceId) // Ask PrinterService to create a print job
        c.JSON(200, iv)
    })
    router.Run(":3000")
}