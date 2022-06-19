package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)


func LoggingSettings(logFile string) {
	f, _ := os.Create("access.log")
    gin.DefaultWriter = io.MultiWriter(f)
}