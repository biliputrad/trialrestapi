package Handler

import (
	"RestAPI-GETNPOST/Entity"
	"RestAPI-GETNPOST/Helper"
	"RestAPI-GETNPOST/Service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type userHandler struct {
	userService Service.UserService
}

func NewUserHandler(userService Service.UserService) *userHandler {
	return &userHandler{userService}
}
func (h *userHandler) GetUserList(c *gin.Context) {
	var usersRespond []Entity.DisplayUser
	_, err, users := h.userService.FindAllUser()
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	for _, u := range users {
		userRespond := Entity.DisplayUser{
			ID:        u.ID,
			UserName:  u.UserName,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Phone:     u.Phone,
		}
		usersRespond = append(usersRespond, userRespond)

	}
	c.JSON(http.StatusOK, gin.H{
		"List of User": usersRespond,
	})

}
func (h *userHandler) GetDataUserById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	_, err, user := h.userService.FindUserByID(int(id))
	final := Helper.ConvertData(user)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": final,
	})
}
func (h *userHandler) LoginHandler(c *gin.Context) {
	var userLogin Entity.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	response, err, user := h.userService.FindUsernameandPassword(userLogin.Username, userLogin.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if response.Success == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			response.Message: "Please Try Again",
		})
	}

	tokenString, err := Helper.GenerateJWT(user.Email, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"Welcome!, here the token": tokenString})
}

func (h *userHandler) RegistrasiHandler(c *gin.Context) {
	var userRequest Entity.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}
	user, err := h.userService.CreateUser(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
func (h *userHandler) UpdateUser(c *gin.Context) {
	var userRequest Entity.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, err := h.userService.UpdateUser(id, userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"Updated Data": user,
	})

}
func (h *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	user, err := h.userService.DeleteUser(int(id))
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Deleteduser": user,
	})
}
func (h *userHandler) ForgetPassword(c *gin.Context) {
	var userForget Entity.ForgetPass
	err := c.ShouldBindJSON(&userForget)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	phone := userForget.Phone
	checkemail, err := h.userService.FindUsernameandEmail(userForget.UserName, userForget.Email)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	if checkemail.Success == true {
		checkphone, err := h.userService.FindUsernameandPhone(userForget.UserName, phone)
		if err != nil {
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
		if checkphone.Success == true {
			_, err := h.userService.ForgetPassword(int(id), userForget)
			if err != nil {
				errorMessages := []string{}
				for _, e := range err.(validator.ValidationErrors) {
					errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
					errorMessages = append(errorMessages, errorMessage)
				}
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": errorMessages,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"Password was changed": userForget.UserName,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				checkphone.Message: "Please Try Again",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			checkemail.Message: "Please Try Again",
		})
	}
}
func (h *userHandler) PaginationUser(c *gin.Context) {

	pagination := Helper.GeneratePaginationRequest(c)

	response, _ := h.userService.PaginationUser(c, pagination)

	if response.Success == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": response.Message,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (h *userHandler) ConvertDataToPDF(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	response, user := h.userService.GenerateUser(id)
	if response.Success == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": response.Message,
		})
		return
	}
	header := []string{"ID", "User Name", "First Name", "Last Name"}
	pdf := Helper.SetToPDF()
	pdf = Helper.Header(pdf, header)
	pdf = Helper.Table(pdf, user)
	if pdf.Err() {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": response.Message,
		})
		return
	}
	err = Helper.SaveFile(pdf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": response.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status ": response.Message,
	})
}
