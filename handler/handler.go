package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mvt-marketplace-backend/repository"
)

type Handler struct {
	repos repository.Repository
}

func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{repos: *repos}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/sale")
	{
		api.POST("/", h.putOnSell)
		api.GET("/", h.getForSaleNftInfo)
	}
	return router
}
