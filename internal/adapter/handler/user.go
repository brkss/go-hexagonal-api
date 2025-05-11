package http

import (
	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/gin-gonic/gin"
)

// UserHandler Represent http-handler for user requests
type UserHandler struct {
	svc port.UserService
}


// NewUserHandler create a new user handller 
func NewUserHandler(svc port.UserService) (*UserHandler){
	return &UserHandler{
		svc,
	}
}

// registerRequest represent the body to register user 
type registerRequest struct {
	Name 		string	`json:"name" binding:"required"`
	Email 		string	`json:"email" binding:"required,email"`
	Password 	string  `json:"password" binding:"required"`
}

func (uh *UserHandler)Register(ctx *gin.Context){

	var req registerRequest;
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ValidationError(ctx, err);
		return 
	}

	user := domain.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}

	_, err := uh.svc.Register(ctx, &user)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)
	HandleSuccess(ctx, rsp) 
}
