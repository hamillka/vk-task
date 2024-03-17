package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"

	"vk-task/internal/handlers/dto"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("auth-x"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			errorDto := &dto.ErrorDto{
				Error: "Токен сформирован неверно",
			}
			err := json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неверный метод подписи: %v", token.Header["alg"])
			}
			return dto.JwtSecretKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			errorDto := &dto.ErrorDto{
				Error: "Неверный токен",
			}
			err = json.NewEncoder(w).Encode(errorDto)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}
