package paseto

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/brkss/dextrace-server/internal/adapter/config"
	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/google/uuid"
)

/*
	PasetoToken implement port.TokenService interface
	and provides an access to the paseto library
*/
type PasetoToken struct {
	token 		*paseto.Token
	key			*paseto.V4SymmetricKey
	parser 		*paseto.Parser
	duration 	time.Duration
}

// new create a new paseto instance
func New(config config.Token) (port.TokenService, error) {
	duration, err := time.ParseDuration(config.Duration)
	if err != nil {
		return nil, domain.ErrTokenDuration;
	}

	token := paseto.NewToken()
	key := paseto.NewV4SymmetricKey()
	parser := paseto.NewParser()

	return &PasetoToken{
		&token,
		&key,
		&parser,
		duration,
	}, nil
}

func (pt *PasetoToken)CreateToken(user *domain.User)(string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &domain.TokenPayload{
		ID: id,
		UserID: user.ID,
	}

	err = pt.token.Set("payload", payload)
	if err != nil {
		return "", domain.ErrTokenDuration
	}

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(pt.duration)

	pt.token.SetIssuedAt(issuedAt)
	pt.token.SetNotBefore(issuedAt)
	pt.token.SetExpiration(expiredAt)

	token := pt.token.V4Encrypt(*pt.key, nil)

	return token, nil
}

func (pt *PasetoToken) VerifyToken(token string) (*domain.TokenPayload, error) {
	var payload *domain.TokenPayload

	parsedToken, err := pt.parser.ParseV4Local(*pt.key, token, nil)
	if err != nil {
		if err.Error() == "this token has expired" {
			return nil, domain.ErrTokenExpired
		}
		return nil, domain.ErrInvalidToken
	}

	err = parsedToken.Get("payload", &payload)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	return payload, nil
}