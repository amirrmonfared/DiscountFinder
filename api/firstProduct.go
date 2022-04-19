package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createFirstProductRequest struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

type createFirstProductRespones struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createFirstProduct(ctx *gin.Context) {
	var req createFirstProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFirstProductParams{
		Brand: req.Brand,
		Link: req.Link,
		Price: req.Price,
	}

	product, err := server.store.CreateFirstProduct(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createFirstProductRespones{
		Brand: product.Brand,
		Link: product.Link,
		Price: product.Price,
		CreatedAt: product.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFirstProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetFirstProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type listProductRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFirstsProduct(ctx *gin.Context) {
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListFirstProductParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	product, err := server.store.ListFirstProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}
