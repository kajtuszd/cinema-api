package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	initializeLogger()
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency)
	})
}

func initializeLogger() {
	t := time.Now()
	timeString := t.Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("logs/cinema-api-%s.log", timeString)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
