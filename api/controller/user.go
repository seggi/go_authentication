package controller

import (
	"fuko_backend/api/service"
	"fuko_backend/models"
	"fuko_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// User controller struct
type UserController struct {
	service service.UserService
}

// New user controller
func FukoUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

// Create user -> calls CreateUser serives for validated user
func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
	}

	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	utils.SuccessJSON(c, http.StatusOK, "Successfully Created user")

}

// Login user: Generates JWT Token for validated user
func (u *UserController) Login(c *gin.Context) {
	var user models.UserLogin
	var hmacSampleSecret []byte
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &utils.Response{
		Success: true,
		Message: "Token generated sucessfully",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}
