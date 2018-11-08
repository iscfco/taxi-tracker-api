package middleware

import (
	"encoding/json"
	"fmt"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/security/jwttasks"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
)

func ValidateToken(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	err, authBackend := jwttasks.NewJwtTasks()
	if err != nil {
		res := errorhandler.HandleErr(&err)
		payload, _ := json.Marshal(res)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(res.HttpCode)
		rw.Write(payload)
		return
	}

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})
	if err != nil || !token.Valid {
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	next(rw, req)
}
