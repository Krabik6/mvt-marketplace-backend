package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getNftId(c *gin.Context) (int, error) {
	id := c.Query("id")
	if len(id) == 0 {
		return 0, errors.New("id not found")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}

func getCollectionAddress(c *gin.Context) (string, error) {
	collection := c.Query("collection")
	if len(collection) == 0 {
		return "", errors.New("collection not correct")
	}

	return collection, nil
}
