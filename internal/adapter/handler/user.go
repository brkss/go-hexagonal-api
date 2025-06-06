package http

import (
	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	Name 		string	`json:"name" binding:"required" example:"John Doe"`
	Email 		string	`json:"email" binding:"required,email" example:"example@example.com"`
	Password 	string  `json:"password" binding:"required" example:"123456789"`
}

func (uh *UserHandler)Register(ctx *gin.Context){

	var req registerRequest;
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ValidationError(ctx, err);
		return 
	}

	uuid, err := uuid.NewRandom();
	if err != nil {
		HandleError(ctx, err);
		return;
	}

	user := domain.User{
		ID: uuid.String(),
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}

	_, err = uh.svc.Register(ctx, &user)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)
	HandleSuccess(ctx, rsp) 
}


func (uh *UserHandler)GetUserProfile(ctx *gin.Context) {
	payload, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		HandleError(ctx, domain.ErrInvalidToken)
		return
	}

	tokenPayload, ok := payload.(*domain.TokenPayload)
	if !ok {
		HandleError(ctx, domain.ErrInvalidToken)
		return
	}

	user, err := uh.svc.GetUser(ctx, tokenPayload.UserID)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	// Remove sensitive data before sending response
	user.Password = ""

	rsp := newUserResponse(*user)
	HandleSuccess(ctx, rsp)
}