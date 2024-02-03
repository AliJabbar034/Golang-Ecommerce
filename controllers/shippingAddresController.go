package controllers

import (
	"github.com/alijabbar034/helper"
	"github.com/alijabbar034/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CReate Shipping Address
func CreateShiipingAddress(c *gin.Context) {

	shiiping := &models.Shipping{}

	if er := c.BindJSON(&shiiping); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, er.Error())
		return
	}

	id, err := shiiping.CreateShipingAddress()

	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	_id, _ := primitive.ObjectIDFromHex(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Created Succesfuly",
		"_id":     _id,
	})
}
