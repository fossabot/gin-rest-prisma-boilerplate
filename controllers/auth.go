// Package controllers includes a collection of controller
//  structures for executing function logic.
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/db"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/forms"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/prisma-client"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/utils"
)

// IAuthenticationController interface.
type IAuthenticationController interface {
	GetSessionData(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

// AuthenticationController for authentication logic.
type AuthenticationController struct{}

var (
	client   = db.DB()
	contextB = context.Background()
)

// GetSessionData method is responsible for retrieving session data once authenticated.
func (AuthenticationController) GetSessionData(c *gin.Context) {
	session := sessions.Default(c)
	uuid := session.Get("uuid")
	if uuid != nil {
		c.JSON(http.StatusOK, gin.H{"uuid": uuid, "username": session.Get("username"),
			"email": session.Get("email")})
	} else {
		utils.CreateError(c, http.StatusBadRequest,
			"Failed to fetch session data. Make sure to be logged in.")
	}
}

// Register method provides registration logic for when onboarding onto the api.
func (AuthenticationController) Register(c *gin.Context) {
	var registrationPayload forms.RegistrationForm
	if validationErr := c.BindJSON(&registrationPayload); validationErr != nil {
		utils.CreateError(c, http.StatusBadRequest, validationErr.Error())
		return
	}
	if hashedPass, hashErr := utils.EncryptPassword(registrationPayload.Password); hashErr != nil {
		utils.CreateError(c, http.StatusInternalServerError, "Failed to hash password.")
	} else {
		registrationPayload.Password = hashedPass
		user, prismaErr := client.CreateUser(prisma.UserCreateInput{
			Email:    registrationPayload.Email,
			Name:     registrationPayload.Name,
			Username: registrationPayload.Username,
			Password: registrationPayload.Password,
			Role:     prisma.RoleDefault,
		}).Exec(contextB)

		if prismaErr != nil {
			log.Print(prismaErr)
			utils.CreateError(c, http.StatusNotAcceptable, "Failed to save profile.")
			return
		}
		// setting session keys
		session := sessions.Default(c)
		session.Set("uuid", user.ID)
		session.Set("email", user.Email)
		session.Set("username", user.Username)
		session.Set("role", string(user.Role))

		if sessionErr := session.Save(); sessionErr != nil {
			utils.CreateError(c, http.StatusInternalServerError, sessionErr.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":     user.Name,
			"username": user.Username,
			"role":     user.Role,
		})
	}
}

// Login method provides login logic for when signing into the api.
func (AuthenticationController) Login(c *gin.Context) {
	var loginPayload forms.LoginForm
	if validationErr := c.BindJSON(&loginPayload); validationErr != nil {
		utils.CreateError(c, http.StatusBadRequest, validationErr.Error())
		return
	}

	if user, err := client.User(
		prisma.UserWhereUniqueInput{Username: &loginPayload.Username}).Exec(contextB); err != nil {
		log.Println(err)
		utils.CreateError(c, http.StatusBadRequest,
			fmt.Sprintf(
				"The profile with the username: %s doesn't exist. Please register before trying to login.",
				loginPayload.Username))
	} else {
		if passwordMatch := utils.CheckPassword(loginPayload.Password, user.Password); passwordMatch != true {
			utils.CreateError(c, http.StatusNotAcceptable, "Invalid password details. Please try again.")
		} else {
			// setting session keys
			session := sessions.Default(c)
			session.Set("uuid", user.ID)
			session.Set("email", user.Email)
			session.Set("username", user.Name)
			session.Set("role", string(user.Role))
			if sessionErr := session.Save(); sessionErr != nil {
				utils.CreateError(c, http.StatusInternalServerError, sessionErr.Error())
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Logged in.",
				"user":    user,
			})
		}
	}
}

// Logout method provides logic for logging out once authenticated.
func (AuthenticationController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if sessionErr := session.Save(); sessionErr != nil {
		log.Print(sessionErr)
		utils.CreateError(c, http.StatusInternalServerError, "Failed to logout.")
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out..."})
}
