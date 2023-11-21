package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mxnuchim/golang-auth-api/initializers"
	"github.com/mxnuchim/golang-auth-api/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	// get data off the request body
	var body struct {
		FirstName string
		LastName  string
		Username  string
		Phone 	  string
		Email     string
		Password  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	//hash password with bcrypt
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	//create user in db
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Phone:     body.Phone,
		Password:  string(bcryptPassword),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{
	        "error": "Error creating user",
	    })
	    return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}