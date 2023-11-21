package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *gin.Context) {
	// get data off request body
	var body struct {
		Email string
		Password string	
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	// find user in db
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
	    c.JSON(http.StatusNotFound, gin.H{
	        "error": "Invalid email or password",
	    })
	    return
	}
	
	// compare password with saved user hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
	    c.JSON(http.StatusUnauthorized, gin.H{
	        "error": "Invalid password",
	    })
	    return
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error signing token",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"name": user.FirstName + " " + user.LastName,
		"email": user.Email,
		"phone": user.Phone,
	})

}