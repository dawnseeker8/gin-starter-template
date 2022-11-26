package handlers

import (
	_ "embed"
	"fmt"
	"net/http"

	"dawnseek.com/gin-starter/core/utils"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
)

// @Title Signin
// @Description sign in as a member
// @Param   code     QueryString    string  true        "The code to sign in"
// @Param   state     QueryString    string  true        "The state"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /signin [post]
// @Tag Account API
func Signin(c *gin.Context) {
	a := ApiHandler{c}
	code := a.Query("code")
	state := a.Query("state")

	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		a.JSON(http.StatusBadRequest, err)
		return
	}

	claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
	if err != nil {
		a.JSON(http.StatusBadRequest, err)
		return
	}

	affected, err := UpdateMemberOnlineStatus(&claims.User, true, utils.GetCurrentTime())
	if affected && err != nil {
		a.JSON(http.StatusBadRequest, err)
		return
	}

	a.SetSessionClaims(claims)
	a.JSON(http.StatusOK, claims)
}

// UpdateMemberOnlineStatus updates member's online information.
func UpdateMemberOnlineStatus(user *casdoorsdk.User, isOnline bool, lastActionDate string) (bool, error) {
	if user == nil {
		return false, fmt.Errorf("user is nil")
	}

	user.IsOnline = isOnline
	SetUserField(user, "lastActionDate", lastActionDate)
	return casdoorsdk.UpdateUserForColumns(user, []string{"isOnline", "properties"})
}

func SetUserField(user *casdoorsdk.User, field string, value string) {
	user.Properties[field] = value
}

// @Title Logout
// @Description sign out the current member
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /Logout [post]
// @Tag Account API
func Logout(c *gin.Context) {
	a := ApiHandler{c}
	claims := a.GetSessionClaims()

	if claims != nil {
		_, err := UpdateMemberOnlineStatus(&claims.User, false, utils.GetCurrentTime())
		if err != nil {
			a.JSON(http.StatusInternalServerError, gin.H{"error": "Logout failure"})
			return
		}

		a.SetSessionClaims(nil)

		a.JSON(http.StatusOK, gin.H{"message": "ok"})
	}

}

// @Title GetAccount
// @Description Get current account
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /get-account [get]
// @Tag Account API
func GetAccount(c *gin.Context) {
	// if c.RequireSignedIn() {
	// 	return
	// }

	a := ApiHandler{c}

	user := a.GetSessionUser()

	if user != nil {
		a.JSON(http.StatusOK, user)
	}
}
