package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response represent a response body format
type response struct {
	Success 	bool	`json:"success" example:"true"`
	Message 	string	`json:"message" example:"Success"`
	Data 		any  	`json:"data,omitempty"`
}

// errorResponse represent a response with an error 
type errorResponse struct {
	Success 	bool `json:"success" example:"false"`
	Message 	[]string `json:"message" example:"Error Message 1, Error Message 2"`
}

// err status map 
var errorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
} 



// NewResponse is helper function that create a new response body 
func NewResponse(success bool, message string, data any) response {
	return response{
		Success: 	success,
		Message: 	message,
		Data: 		data,
	}
}

// NewErrorResponse return new error response body 
func NewErrorResponse(errs []string) (errorResponse) {
	return errorResponse{
		Success: false,
		Message: errs,
	}
}

// VaidationError send response to some request validation specific errors 
func ValidationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err);
	errResponse := NewErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errResponse)
}

// HandleError determine error code status and return a JSON response with error message and status code 
func HandleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsgs := parseError(err)
	errResponse := NewErrorResponse(errMsgs)
	ctx.JSON(statusCode, errResponse)
}

// parseError parse error messages from error object and return them into a slice 
func parseError(errs error) [] string {
	var errMsgs []string;

	if errors.As(errs, &validator.ValidationErrors{}) {
		for _, err := range errs.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	}else {
		errMsgs = append(errMsgs, errs.Error())
	}

	return errMsgs;
}


type userResponse struct {
	ID 		string `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
	Password	string `json:"password"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

func newUserResponse(user domain.User)  userResponse{
	return userResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// HandleSuccess handle successfuly response 
func HandleSuccess(ctx *gin.Context, data any) {
	rsp := NewResponse(true, "success", data)
	ctx.JSON(http.StatusOK, rsp)
}

// authResponse represent an authentication response body 
type authResponse struct {
	AccessToken string `json:"token"`
}

// newAuthResponse create a new auth response body 
func newAuthResponse(token string) authResponse {
	return  authResponse {
		AccessToken: token,
	}
}

