package controllers

import (
	"fmt"
	"github.com/alijabbar034/helper"
	"github.com/alijabbar034/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"path/filepath"
	"strconv"
)

// Create Product
func CreateProduct(c *gin.Context) {
	product := &models.Product{}
	reqUser, _ := c.Get("user")
	user := reqUser.(*models.User)
	name := c.PostForm("name")
	product.Description = c.PostForm("description")
	reqPrice := c.PostForm("price")
	price, _ := strconv.ParseFloat(reqPrice, 64)
	product.Price = price
	product.CreatedBy = user.ID
	product.Color = c.PostFormArray("color")

	files, ok := c.Request.MultipartForm.File["file"]
	if !ok {
		helper.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}
	uploadPath := "./resources/products/images"
	var fileName []string
	for _, file := range files {
		filename := fmt.Sprintf("%s_%s", name, file.Filename)
		filePath := filepath.Join(uploadPath, filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			helper.ErrorHandler(c, http.StatusInternalServerError, "Erro during uploading files")
			return
		}
		fileName = append(fileName, filePath)
	}

	if name != "" {
		product.Name = name
	}
	product.Images = fileName

	id, eror := product.CreateProduct()
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Eror during creation of product")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Created Product Successfully",
		"productId": fileName,
		"id":        id,
	})

}

// Get All Prodct
func GetAllProduct(c *gin.Context) {
	query := bson.D{}
	lessPrice := c.Query("lt")
	greatePrice := c.Query("gt")
	color := c.Query("color")
	if lessPrice != "" {
		less, _ := strconv.Atoi(lessPrice)
		query = append(query, bson.E{"price", bson.D{
			{"$lt", less},
		}})
	}
	if greatePrice != "" {
		greate, _ := strconv.Atoi(greatePrice)
		query = append(query, bson.E{"price", bson.D{{
			"$gt", greate},
		}})
	}
	if color != "" {
		query = append(query, bson.E{"color", color})
	}
	products, err := models.GetProducts(query)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "all products",
		"products": products,
	})
}

//Get Product By Id

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	product, er := models.GetByID(id)
	if er != nil {
		helper.ErrorHandler(c, http.StatusNotFound, "No Product Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get product by id",
		"product": product,
	})
}

//Update Product

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if er := c.BindJSON(&product); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, er.Error())
		return
	}
	_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	updated, eror := models.UpdateProduct(product, _id)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "updated Successfulyy",
		"count":   updated,
	})
}

//Delete Product

func DeleteProduct(c *gin.Context) {
	_id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(_id)

	deletedCount, eror := models.DeleteProduct(id)
	if eror != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, eror.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Delteed",
		"count":   deletedCount,
	})
}
