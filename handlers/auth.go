package handlers

import (
	"net/http"

	"woyteck.pl/gothstarter/views/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
