package middlewares

import (
	"ff/api/controllers"
	"net/http"
)

func RequireRole(requiredRoles []string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := controllers.GetCurrentLoggedInUser(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		hasRole := false
		for _, role := range requiredRoles {
			if role == user.Role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			http.Error(w, "Insufficient permissions", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
