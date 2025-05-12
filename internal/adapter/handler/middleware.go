package http

import (
	"strings"

	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/gin-gonic/gin"
)



const (
	// authorizationKey is the key for authorization header in the request 
	authorizationKey = "authorization"
	// authorizationType is the accespted authorization type by the app 
	authorizationType = "bearer"
	// authorizationPayloadKey is the token payload key in the context 
	authorizationPayloadKey = "authorization_payload"
)

// authorizationMiddleware is the middleware that checks if the user id authenticated 
func authorizationMiddleware(token port.TokenService) gin.HandlerFunc {
	return func (ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationKey)
		isEmpty := len(authorizationHeader) == 0
		if isEmpty {
			err := domain.ErrEmptyAuthorizationHeader
			handleAbort(ctx, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := domain.ErrInvalidAuthorizationHeader
			handleAbort(ctx, err)
			return
		} 

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != authorizationType {
			err := domain.ErrInvalidAuthorizationType
			handleAbort(ctx, err)
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			handleAbort(ctx, err)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}