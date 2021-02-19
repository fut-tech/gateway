package middleware

import "github.com/gin-gonic/gin"

// AuthHandler Authorization handler
type AuthHandler struct {
}

// NewAuthHandler Make new struct
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// CheckToken Check token handler
func (a *AuthHandler) CheckToken(ctx *gin.Context) {
	var token, err = ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}

	if token != "test" {
		ctx.AbortWithStatus(401)
		return
	}

	ctx.Next()
}
