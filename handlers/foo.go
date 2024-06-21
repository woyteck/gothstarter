package handlers

import (
	"net/http"

	"woyteck.pl/gothstarter/views/foo"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return foo.Index().Render(r.Context(), w)
}
