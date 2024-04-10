package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// issuer is the issuer of the jwt token.
	Issuer = "travegram"
	// Signing key section. For now, this is only used for signing, not for verifying since we only
	// have 1 version. But it will be used to maintain backward compatibility if we change the signing mechanism.
	KeyID = "v1"
	// AccessTokenAudienceName is the audience name of the access token.
	AccessTokenAudienceName = "user.access-token"
	AccessTokenDuration     = 7 * 24 * time.Hour

	// CookieExpDuration expires slightly earlier than the jwt expiration. Client would be logged out if the user
	// cookie expires, thus the client would always logout first before attempting to make a request with the expired jwt.
	CookieExpDuration = AccessTokenDuration - 1*time.Minute
	// AccessTokenCookieName is the cookie name of access token.
	AccessTokenCookieName = "travegram.access-token"
)

type ClaimsMessage struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string, userID int32, accessTokenExpirationTime time.Time) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		Issuer:    Issuer,
		Audience:  jwt.ClaimStrings{AccessTokenAudienceName},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprint(userID),
		ExpiresAt: jwt.NewNumericDate(accessTokenExpirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ClaimsMessage{
		Name:             username,
		RegisteredClaims: registeredClaims,
	})

	tokenString, err := token.SignedString([]byte("travegram"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
