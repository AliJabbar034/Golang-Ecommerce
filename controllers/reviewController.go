package controllers

import (
	"github.com/alijabbar034/helper"
	"github.com/alijabbar034/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// Create Review
func CreateReview(c *gin.Context) {

	review := &models.Review{}
	id := c.Param("id")
	productId, _ := primitive.ObjectIDFromHex(id)
	if err := c.BindJSON(&review); err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Invalid data")
		return
	}
	review.ProductId = productId

	id, er := review.CreateReview()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Error during creating review")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "review added succesffuly",
		"id":      id,
	})
}

// Delete REview
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ErrorHandler(c, http.StatusBadRequest, "Plaese Provide Id")
		return
	}
	_id, _ := primitive.ObjectIDFromHex(id)

	deletedCount, err := models.Delete(_id)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"count": deletedCount,
	})
}

// Get All Review
func GetAllReviews(c *gin.Context) {

	reviews, er := models.GetReviews()
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total":   len(reviews),
		"reviews": reviews,
	})
}

// Updated REview
func UpdateReview(c *gin.Context) {

	var review models.Review
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)
	if err := c.BindJSON(&review); err != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, "Provide the valid data")
		return
	}

	updated, er := models.UpdateRev(review, id)
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"updatedCount": updated,
	})

}

//Get REview By Id

func GetRevById(c *gin.Context) {
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	review, er := models.GetRev(id)
	if er != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, er.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Got succesfuly",
		"review":  review,
	})
}
