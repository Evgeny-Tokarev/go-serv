package helper

import (
	"errors"
	"fmt"
	"github.com/Evgeny-Tokarev/go-serv/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strconv"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type MyClaims struct {
	ID  int64 `json:"id"`
	IAT int64 `json:"iat"`
	EAT int64 `json:"eat"`
}

func (claims *MyClaims) Valid() error {
	now := time.Now().Unix()
	if now > claims.EAT {
		return jwt.NewValidationError("Token has expired", jwt.ValidationErrorExpired)
	}
	return nil
}

func GenerateJwtCookie(user model.User) (*http.Cookie, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Second * time.Duration(tokenTTL)),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie, nil
}

func CurrentUser(context *gin.Context) (model.User, error) {
	err := ValidateJWT(context)
	if err != nil {
		return model.User{}, err
	}
	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := model.FindUserById(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	cookie, err := context.Request.Cookie("jwt")
	if err != nil {
		return nil, fmt.Errorf("jwt cookie not found: %v", err)
	}

	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	if err != nil {
		// log the error or provide more context in the error message
		return nil, fmt.Errorf("failed to parse JWT token: %v", err)
	}
	return token, err
}

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		fmt.Println("no token: ", err)
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}
