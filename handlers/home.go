package handlers

import (
	"net/http"

	"woyteck.pl/gothstarter/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
