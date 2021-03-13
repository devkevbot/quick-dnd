package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *application) getJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey:  []byte(app.jwtSigningKey),
		ContextKey:  "token",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}
}

func (app *application) createJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(app.jwtSigningKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getUsernameFromToken(c echo.Context) string {
	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	tokenUsername := claims["username"].(string)
	return tokenUsername
}
