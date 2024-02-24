package middleware

import (
	"net/http"
	"strings"
	"synapsis/src/config"
	"synapsis/src/utils"

	"golang.org/x/net/context"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		if reqToken == "" {
			utils.JSONError(w, "Invalid Token", http.StatusBadRequest)
			return
		}

		config, err := config.LoadConfig(".")
		if err != nil {
			utils.JSONError(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		id, err := utils.ValidateToken(reqToken, []byte(config.AccessKey))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "ID", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
