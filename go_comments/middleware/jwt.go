package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// ValidateJWT verifies the JWT token using the DRF service
func ValidateJWT(next http.Handler, tokenVerifyURL string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"error": "Unauthorized: Missing token"}`, http.StatusUnauthorized)
			return
		}

		// Extract the token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token with the DRF service
		req, err := http.NewRequest("POST", tokenVerifyURL, strings.NewReader(fmt.Sprintf(`{"token": "%s"}`, token)))
		if err != nil {
			http.Error(w, `{"error": "Failed to create verification request"}`, http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, `{"error": "Unauthorized: Invalid token"}`, http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
