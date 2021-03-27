package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func IsAuthorized(next http.Handler) http.Handler {
  var mySigningKey = []byte(os.Getenv("API_KEY"))

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Header["Authorization"] != nil {
      tokenSplited := strings.SplitAfter(r.Header["Authorization"][0], " ")[1]      
      token, err := jwt.Parse(tokenSplited, func(token *jwt.Token) (interface{}, error) {
          if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("There was an error")
          }
          return mySigningKey, nil
      })

      if err != nil {
        ResponseWithError(w, http.StatusInternalServerError, "Unauthorized")
      }

      if token.Valid {
        next.ServeHTTP(w, r)
      }
    } else {
      ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
    }
  })
}