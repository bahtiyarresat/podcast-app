package gateway

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

const (
	allowedRequests   = 5
	secondsPerRequest = 1
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {

	router := gin.Default()
	router.Use(rateLimitMiddleware(allowedRequests, secondsPerRequest))
	server := &Server{router}
	router.GET("/podcasts", server.getAllPodcasts)
	router.GET("/heartbeat", heartbeat)

	return server
}
func heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func rateLimitMiddleware(requests int, durationSeconds int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(requests), durationSeconds)

	return func(c *gin.Context) {
		if limiter.Allow() == false {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// start server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
