package routes

import (
	"github.com/alijabbar034/controllers"
	"github.com/alijabbar034/middleware"
	"github.com/gin-gonic/gin"
)

func ReviewRouter(c *gin.RouterGroup) {
	router := c.Group("/review")
	router.POST("/:id", middleware.Authenticate, controllers.CreateReview)
	router.DELETE("/:id", middleware.Authenticate, controllers.DeleteReview)
	router.GET("/", controllers.GetAllReviews)
	router.PUT("/:id", middleware.Authenticate, controllers.UpdateReview)
	router.GET("/:id")
}
