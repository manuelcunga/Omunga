package middlewares

import (
	"net/http"
)

func MyProtectedHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	// Use the user ID to retrieve data or perform other actions
}
