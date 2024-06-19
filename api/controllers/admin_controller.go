package controllers

import (
	"net/http"
)

func AdminUpdateUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userUUID := r.FormValue("userUUID")
	role := r.FormValue("role")

	if userUUID == "" || role == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	UpdateRoleByUserUUID(userUUID, role)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
