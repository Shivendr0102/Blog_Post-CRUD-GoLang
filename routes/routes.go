package routes

import (
	"22nd_Oct_Antino/blog/posts"
	"22nd_Oct_Antino/config"
	"22nd_Oct_Antino/db"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func New(configuration *config.Config, dbSource *db.Source) *chi.Mux {

	router := chi.NewRouter()
	corsPolicy := cors.New(cors.Options{
		AllowedOrigins:   configuration.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token", "access-control-allow-origin", "access-control-expose-headers", "content-type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		Debug:            false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(
		corsPolicy.Handler, // Allow access to the API from everywhere
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
	)
	router.Route("/blog", func(r chi.Router) {
		r.Group(func(r chi.Router) {

			r.Mount("/blogposts/", posts.NewModule(configuration, dbSource).Routes())

			// TODO Later For Each Blog / User

			// r.Mount("/comments/", comments.NewModule(configuration, dbSource).Routes())
			// r.Mount("/reactions/", comments.NewModule(configuration, dbSource).Routes())
			// r.mount("/similarposts/", comments.NewModule(configuration, dbSource).Routes())

		})
	})

	return router
}
