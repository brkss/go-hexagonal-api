package http

import (
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/gin-gonic/gin"
)

// AuthHandler represent the HTTP-handler for authentication related requests
type AuthHandler struct {
	port.AuthService
}


func NewAuthHandler(svc port.AuthService)(*AuthHandler){
	return &AuthHandler{
		svc,
	}
}

type loginRequest struct {
	Email 		string `json:"email" binding:"required,email" example:"example@example.com"`
	Password 	string `json:"password" binding:"required" example:"123456789"`
}

func (ah *AuthHandler)Login(ctx *gin.Context) {
	
	var req loginRequest;
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ValidationError(ctx, err)
		return
	}

	
	
	token, err := ah.AuthService.Login(ctx, req.Email, req.Password)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	rsp := newAuthResponse(token)

	HandleSuccess(ctx, rsp)
	
}


