package middlewares

import (
	"auth-api/connections"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase-community/supabase-go"
)

type contexTokenKey string

var userIdKey = contexTokenKey("user_id")

func AuthMiddleware(cli *supabase.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		auth := req.Header.Get("Authorization")

		if auth == "" {
			http.Error(res, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		jwtSecret := os.Getenv("SUPA_TOKEN")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(res, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["sub"].(string)

		ctx := context.WithValue(req.Context(), userIdKey, userId)

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func GetDataLogin(res http.ResponseWriter, req *http.Request) {

	var userRol struct {
		Rol string `json:"rol"`
	}

	userId := req.Context().Value(userIdKey)

	if userId == nil {
		http.Error(res, "Unauthorized: missing user ID", http.StatusUnauthorized)
		return
	}

	userIdV, ok := userId.(string)

	if !ok {
		http.Error(res, "Invalid user ID type", http.StatusInternalServerError)
		return
	}

	_, err := connections.Client.From("readers").Select("rol", "", false).Eq("id", userIdV).Single().ExecuteTo(&userRol)

	if err != nil {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(res).Encode(userRol)
}

func CheckRole(cli *supabase.Client, userId string, requiried string) bool {
	var clientRole struct {
		Role string `json:"role"`
	}

	_, err := connections.Client.From("readers").Select("role", "", false).Eq("id", userId).Single().ExecuteTo(&clientRole)

	if err != nil {
		return false
	}

	return clientRole.Role == requiried
}

func AdminAuthMiddleware(cli *supabase.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		userId := req.Context().Value(userIdKey).(string)

		if !CheckRole(cli, userId, "admin") {
			http.Error(res, "Access not authorized", http.StatusForbidden)
		}

		next.ServeHTTP(res, req)
	})
}
