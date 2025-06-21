package web

import (
	inertia "github.com/romsar/gonertia"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RootHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Home/Index", inertia.Props{
			"text": "Inertia.js with React and Go! 💚",
		})
		if err != nil {
			log.Error().Err(err).Msg("❌ Render failed")
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	})
}
