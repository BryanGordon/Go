package middlewares

import (
	"net/http"

	"github.com/supabase-community/supabase-go"
)

func AuthMiddleware(cli *supabase.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		auth := ""
	})
}
