package middleware

import (
    "context"
    "net/http"

    "github.com/go-chi/jwtauth/v5"
    "github.com/gasparvini/configs"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verifica o token JWT a partir do cabeçalho Authorization
        token, claims, err := jwtauth.FromContext(r.Context())
        if err != nil || token == nil || !jwtauth.IsValid(token) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Opcional: adicionar informações ao contexto, como os claims do JWT
        ctx := context.WithValue(r.Context(), "userClaims", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}