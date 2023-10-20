package middleware

import (
	"net/http"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(next http.Handler, jwtSecretKey string) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("jwt")
        if err != nil || cookie.Value == "" {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecretKey), nil
        })

        if err != nil || !token.Valid {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        next.ServeHTTP(w, r)
    })
}
