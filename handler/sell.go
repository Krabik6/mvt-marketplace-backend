package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mvt-marketplace-backend/models"
	"net/http"
)

func (h *Handler) putOnSell(ctx *gin.Context) {
	var input models.SellInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repos.PutOnSell(input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) getForSaleNftInfo(ctx *gin.Context) {
	id, err := getNftId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection, err := getCollectionAddress(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input = models.GetNftInfoInput{
		NftId:         id,
		NftCollection: collection,
	}

	fmt.Println(input)

	output, err := h.repos.GetForSaleNftInfo(input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
