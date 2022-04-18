package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("bebas")
var tokenName = "token"

type Claims struct {
	ID       int    `json:id`
	Name     string `json :"name"`
	UserType string `json:user_type`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, name string, userType string) {

	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{

		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{

		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})

}

func resetUserToken(w http.ResponseWriter) {

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})

}

func Authenticate(next http.HandlerFunc, accessType string) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(w, r, accessType)
		if !isValidToken {
			sendUnAuthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})

}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType string) bool {

	isAccessTokenValid, id, email, userType := validateTokenFromCookies(r)
	fmt.Print(id, email, userType, accessType, isAccessTokenValid)
	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false

}

func validateTokenFromCookies(r *http.Request) (bool, int, string, string) {

	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parseToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parseToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}
	}

	return false, -1, "", ""

}
