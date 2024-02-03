package controllers

import (
	"fmt"
	"github.com/alijabbar034/helper"
	"github.com/alijabbar034/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Register User

func RegisterUser(c *gin.Context) {
	user := &models.User{}
	if eror := c.ShouldBindJSON(&user); eror != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	if user.FName == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	hashPass, hashError := helper.HashPassword(user.Password)
	if hashError != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Error during Hashing passsword")
		return
	}
	user.Password = hashPass
	user.Role = "user"
	id, err := user.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	helper.SendToken(c, id, "user created Successfully")

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "User Created succesfully",
	//	"id":      id,
	//})
}

//Login User

func LoginUser(c *gin.Context) {

	var user models.User
	if eror := c.BindJSON(&user); eror != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, "Inavalid data")
		return
	}
	if user.Password == "" || user.Email == "" {
		helper.ErrorHandler(c, http.StatusBadRequest, "please provide email and password")
		return
	}
	foundUser, err := models.UserLogin(user.Email)
	if err != nil {
		helper.ErrorHandler(c, http.StatusUnauthorized, "No user Found ")
		return
	}
	fmt.Println("hhhhhhhh", foundUser)
	compareEror := helper.CompareHashedPassword(user.Password, foundUser.Password)
	if compareEror != nil {
		helper.ErrorHandler(c, http.StatusUnauthorized, "wrong password")
		return
	}
	_id := foundUser.ID.Hex()
	helper.SendToken(c, _id, "Login Successfully")
}

// Get User By Id
func GetUser(c *gin.Context) {

	id := c.Param("id")

	user, err := models.GetById(id)
	if err != nil {
		helper.ErrorHandler(c, http.StatusNotFound, "No user Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

//Get All User

func GetAllUsers(c *gin.Context) {

	users, err := models.GetAllUser()
	if err != nil {
		helper.ErrorHandler(c, http.StatusNotFound, "No users Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "all user list",
		"total":   len(users),
		"users":   users,
	})
}

//UPadte User Profile

func UpdateUser(c *gin.Context) {

	reqUser, _ := c.Get("user")
	userObj, _ := reqUser.(*models.User)
	var user models.User
	if er := c.BindJSON(&user); er != nil {
		helper.ErrorHandler(c, http.StatusBadRequest, "Inavlid info...")
		return
	}
	//if user.Email == "" || user.Password == "" || user.FName == "" {
	//	helper.ErrorHandler(c, http.StatusBadRequest, "Provide all info")
	//	return
	//}

	id, err := models.UpdateProfile(user, userObj.ID)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "Error  during updating user")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated Sucessfuly",
		"count":   id,
	})
}

// LOgout
func LogoutUser(c *gin.Context) {
	c.SetCookie("token", "", 1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout succesfully",
	})
}

// GEt Profile
func GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		helper.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
