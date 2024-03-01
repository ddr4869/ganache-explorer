package internal

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ddr4869/ether-go/config"
	"github.com/ddr4869/ether-go/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	router    *gin.Engine
	config    *config.Config
	entClient *ent.Client
}

func NewRestController(cfg *config.Config) (*Server, error) {
	gin.SetMode(cfg.Gin.Mode)
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(GetGinLogFomatter()))
	router.Use(corsMiddleware())
	entClient, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=1821 sslmode=disable")
	if err != nil {
		return nil, err
	}
	// defer entClient.Close()
	// Run the auto migration tool.
	if err := entClient.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	server := &Server{router, cfg, entClient}
	SetUp(server)
	return server, nil
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}

	config.AddAllowHeaders("Authorization")
	config.AddAllowHeaders("x-frame-options")
	config.AddAllowHeaders("Cache-Control")
	config.AddAllowHeaders("X-XSS-Protection")
	config.AddAllowHeaders("Referrer-Policy")
	config.AddAllowHeaders("Content-Security-Policy")
	config.AddAllowHeaders("Feature-Policy")

	return cors.New(config)
}

func GetGinLogFomatter() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor, ginColor string
		var ginColorInt int
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
			ginColorInt, _ = strconv.Atoi(param.StatusCodeColor()[5:7])
			ginColor = fmt.Sprintf("\033[%dm", ginColorInt-10)
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}

		return fmt.Sprintf("%sGIN%s    [%s] |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			ginColor, resetColor,
			param.TimeStamp.Format(time.RFC3339),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}
}
