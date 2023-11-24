package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var mySigningKey = []byte(DotEnvVariable("JWT_SECRET"))

// IsAuthorized -> verify jwt header
func IsAuthorized(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenHeader := r.Header.Get("Token")
		if tokenHeader != "" {
			token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				AuthorizationResponse("Invalid JWT token", w)
			}

			if token.Valid {
				next(w, r, ps)
			}
		} else {
			AuthorizationResponse("Not Authorized", w)
		}
	}
}

/*
func isAuthorzedFirebase(next httprouter.Handle) httprouter.Handle {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
}
*/
// GenerateJWT -> generate jwt
func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
