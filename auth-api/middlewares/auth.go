package middlewares

import (
	"auth-api/connections"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase-community/supabase-go"
)

func AuthMiddleware(cli *supabase.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		type contexTokenKey string
		var userIdKey = contexTokenKey("user_id")

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
		userId := req.Context().Value("user_id").(string)

		if !CheckRole(cli, userId, "admin") {
			http.Error(res, "Access not authorized", http.StatusForbidden)
		}

		next.ServeHTTP(res, req)
	})
}
