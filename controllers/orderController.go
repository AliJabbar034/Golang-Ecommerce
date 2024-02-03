package controllers

import (
	"github.com/alijabbar034/helper"
	"github.com/alijabbar034/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// Create Order
func CreateOrder(c *gin.Context) {

	order := &models.Order{}
	if err := c.BindJSON(&order); err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, er := order.CreateOrder()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created Successfully",
		"id":      id,
	})

}

// Get All Order
func GetAllOrders(c *gin.Context) {

	orders, er := models.GetAllOrder()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all order",
		"orders":  orders,
	})
}

// Get Order By Id
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	_id, _ := primitive.ObjectIDFromHex(id)
	order, eror := models.GetAOrder(_id)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"order":   order,
	})
}

//Update Order Status

func UpdateOrde(c *gin.Context) {
	var order models.Order

	if er := c.BindJSON(&order); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, er.Error())
		return
	}
	if order.Status == "" {
		helper.ErrorHandler(c, http.StatusBadRequest, "Order Status is empty")
		return
	}
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	count, eror := models.UpdateOrderStatus(id, order.Status)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated successfuly",
		"count":   count,
	})

}
