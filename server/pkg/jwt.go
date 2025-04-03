package pkg

import (
	"errors"
	"fmt"
	"os"
	"time"

	"absensi/internal/entity"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	User entity.User `json:"user"`
}

type RefreshClaims struct {
	jwt.RegisteredClaims
}

var (
	jwtAccessKeyExpirationTime  = time.Now().Add(5 * time.Minute)
	jwtRefreshKeyExpirationTime = time.Now().Add(24 * time.Hour)
)

var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

func JWTEncode(user entity.User) (*entity.UserJWT, error) {
	// Create the JWT claims, which includes the user data and expiry time
	claims := &UserClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwtAccessKeyExpirationTime),
		},
	}
	refreshClaims := &RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwtRefreshKeyExpirationTime),
		},
	}

	// Declare the tokenAccessKey with the algorithm used for signing, and the claims
	tokenAccessKey := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create access key
	tokenAccessKeyStr, err := tokenAccessKey.SignedString(jwtSecretKey)
	if err != nil {
		return nil, err
	}
	tokenRefreshKey := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenRefreshKeyStr, err := tokenRefreshKey.SignedString(jwtSecretKey)
	if err != nil {
		return nil, err
	}
	userJWT := entity.UserJWT{
		AccessToken:  tokenAccessKeyStr,
		RefreshToken: tokenRefreshKeyStr,
	}
	return &userJWT, nil
}

func JWTDecode(s string) (*entity.User, error) {
	userClaims := &UserClaims{}
	token, err := jwt.ParseWithClaims(s, userClaims, func(token *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	})
	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return &userClaims.User, nil

}

func JWTRefresh(accessKey string, refreshKey string) (*entity.User, *entity.UserJWT, error) {
	// refresh key
	refreshClaims := &RefreshClaims{}
	tknRefreshKey, err := jwt.ParseWithClaims(refreshKey, refreshClaims, func(token *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	if !tknRefreshKey.Valid {
		return nil, nil, ErrUnauthorized // not valid
	}
	// If refresh token already timeout return err
	if time.Now().After(refreshClaims.ExpiresAt.Time) {
		return nil, nil, ErrUnauthorized // not enough time to refresh
	}
	refreshClaims.ExpiresAt = jwt.NewNumericDate(jwtRefreshKeyExpirationTime)
	tokenRefreshKey := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	tokenRefreshKeyStr, err := tokenRefreshKey.SignedString(jwtSecretKey)
	if err != nil {
		return nil, nil, err
	}
	// refresh key
	accessClaims := &UserClaims{}
	tknAccessKey, err := jwt.ParseWithClaims(accessKey, accessClaims, func(token *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	})
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, nil, err
	} else {
		tknAccessKey.Valid = true // LOLLLL
	}
	if !tknAccessKey.Valid {
		return nil, nil, ErrUnauthorized // not valid
	}
	accessClaims.ExpiresAt = jwt.NewNumericDate(jwtAccessKeyExpirationTime)
	tokenAccessKey := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	tokenAccessKeyStr, err := tokenAccessKey.SignedString(jwtSecretKey)
	if err != nil {
		return nil, nil, err
	}

	userJWT := entity.UserJWT{
		AccessToken:  tokenAccessKeyStr,
		RefreshToken: tokenRefreshKeyStr,
	}
	return &accessClaims.User, &userJWT, nil
}
