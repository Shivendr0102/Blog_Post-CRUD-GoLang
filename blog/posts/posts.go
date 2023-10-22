//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml

package posts

import (
	"22nd_Oct_Antino/config"
	"22nd_Oct_Antino/db"

	"github.com/go-chi/chi"
)

type Config struct {
	*config.Config
	dbSource *db.Source
}

func NewModule(configuration *config.Config, dbSource *db.Source) *Config {
	return &Config{configuration, dbSource}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()

	// Creating routes for each API call under Admin package

	// BlogPost Module APIs
	router.Post("/addBlogPost", config.AddBlogPost)
	router.Put("/updateBlogPostById", config.UpdateBlogPostById)
	router.Delete("/deleteBlogPostById", config.DeleteBlogPostById)
	router.Post("/getallBlogPosts", config.GetAllBlogPosts)
	router.Get("/getBlogPostbyid", config.GetBlogPostById)

	return router
}
