package api

import (
	db "github.com/amirrmonfared/WebCrawler/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for our scraper service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/Products", server.createProduct)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
