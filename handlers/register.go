package handlers

import (
	"fmt"

	"net/http"

	"github.com/geoairng/Finalyear/model"
	"github.com/geoairng/Finalyear/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body schemas.RegisterInput true "Registration Request"
// @Success 201 {object} schemas.Success "User registered successfully"
// @Failure 400 {object}  schemas.OutputMessage "Errror: Bad Request"
// @Failure 409 {object}  schemas.OutputMessage "Errror: Conflict"
// @Router /user [post]
func (s *Server) Register(c *gin.Context) {

	var input schemas.RegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(400, schemas.OutputMessage{
			ErrorMessage: err.Error()})
		c.Abort()
		return
	}

	real := input.VerifyPassword()

	if !real {
		c.JSON(400, schemas.OutputMessage{ErrorMessage: "Both the password field and confirm password field doesn't match"})
		c.Abort()
		return
	}

	real = input.ValidateEmail()

	if !real {
		c.JSON(400, schemas.OutputMessage{ErrorMessage: "Email doesn't match the correct pattern"})
		c.Abort()
		return
	}

	err = input.HashPassword()

	if err != nil {
		fmt.Println(err.Error())
	}

	valid := model.User{
		CompanyName: input.CompanyName,
		Email:       input.Email,
		Password:    input.Password,
	}

	result := s.DB.Create(&valid)

	if result.Error != nil {
		c.JSON(http.StatusConflict, schemas.OutputMessage{ErrorMessage: "Email or Username has already exist"})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, schemas.Success{Message: "User Created"})

}
