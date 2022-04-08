package api

import (
	"net/http"

	db "github.com/amirrmonfared/WebCrawler/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createProductRequest struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string  `json:"price"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFirstParams{
		Brand: req.Brand,
		Link:  req.Link,
		Price: req.Price,
	}

	account, err := server.store.CreateFirst(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
