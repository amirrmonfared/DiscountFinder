package api

import (
	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for our scraper service.
type Server struct {
	store  db.Store
	router *gin.Engine
	config util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
	}
	router := gin.Default()

	router.POST("/product", server.createProduct)
	router.GET("/product/:id", server.getProduct)
	router.GET("/products", server.listProduct)

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
