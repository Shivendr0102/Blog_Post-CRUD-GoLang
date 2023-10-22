package posts

import (
	"22nd_Oct_Antino/db/blog"
	"22nd_Oct_Antino/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

func ErrorResponse(r *http.Request, w http.ResponseWriter, message string, httpStatus int) {
	render.Status(r, httpStatus)
	render.JSON(w, r, model.Error{
		Message:   message,
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
	})
}

// swagger:route POST /BlogPost
// Returns BlogPost Id
// responses:
//
//	  200:
//		 description: 'OK'
//
// AddBlogPost returns blogpostId in response

func (config *Config) AddBlogPost(w http.ResponseWriter, r *http.Request) {

	var blog_post model.BlogPost
	err := json.NewDecoder(r.Body).Decode(&blog_post)

	// log.Error(err)
	if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if blog_post.Title == nil || strings.TrimSpace(*blog_post.Title) == "" {
		ErrorResponse(r, w, "Please provide Title.", http.StatusBadRequest)
		return
	}

	if blog_post.Body == nil || strings.TrimSpace(*blog_post.Body) == "" {
		ErrorResponse(r, w, "There is not content in Body.", http.StatusBadRequest)
		return
	}

	if blog_post.First_Name == nil || strings.TrimSpace(*blog_post.First_Name) == "" {
		ErrorResponse(r, w, "Please provide post writer`s name.", http.StatusBadRequest)
		return
	}

	blogPostId, err := config.dbSource.Blog.AddBlogPost(blog_post)
	if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, blogPostId)
}

// swagger:route PUT /BlogPost
// responses:
//   200:
// 	 description: 'OK'
// UpdateBlogPostById returns number of rows affected

func (config *Config) UpdateBlogPostById(w http.ResponseWriter, r *http.Request) {

	var blog_post model.BlogPost
	err := json.NewDecoder(r.Body).Decode(&blog_post)

	if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if blog_post.Title == nil || strings.TrimSpace(*blog_post.Title) == "" {
		ErrorResponse(r, w, "Please provide Title.", http.StatusBadRequest)
		return
	}

	if blog_post.Body == nil || strings.TrimSpace(*blog_post.Body) == "" {
		ErrorResponse(r, w, "There is not content in Body.", http.StatusBadRequest)
		return
	}

	rowsAffected, err1 := config.dbSource.Blog.UpdateBlogPostById(blog_post)
	if err1 == blog.ErrPostNotFound {
		log.Error(err1)
		ErrorResponse(r, w, err1.Error(), http.StatusNotFound)
		return
	} else if err1 != nil {
		log.Error(err1)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, rowsAffected)
}

// swagger:route DELETE /BlogPost
// responses:
//
//	  200:
//		 description: 'OK'
//
// DeleteBlogPostById - no return value

func (config *Config) DeleteBlogPostById(w http.ResponseWriter, r *http.Request) {
	blogPostId, err := strconv.ParseInt(r.URL.Query().Get("blogPostId"), 10, 64)
	if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err1 := config.dbSource.Blog.DeleteBlogPostById(blogPostId)
	if err1 == blog.ErrPostNotFound {
		log.Error(err1)
		ErrorResponse(r, w, err1.Error(), http.StatusNotFound)
		return
	} else if err1 != nil {
		log.Error(err1)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
}

// swagger:route GET /BlogPost
// responses:
//   200:
// 	 description: 'OK'
// GetAllBlogPosts - list of all BlogPosts

func (config *Config) GetAllBlogPosts(w http.ResponseWriter, r *http.Request) {

	var inputModel model.SearchPostRequest
	err1 := json.NewDecoder(r.Body).Decode(&inputModel)
	if err1 != nil {
		log.Error(err1)
		ErrorResponse(r, w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	blogPostList, err := config.dbSource.Blog.GetAllBlogPosts(&inputModel)
	if err == blog.ErrPostNotFound {
		log.Error(err)
		ErrorResponse(r, w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, blogPostList)
	render.Status(r, http.StatusOK)
}

// swagger:route GET /BlogPost
// responses:
//   200:
// 	 description: 'OK'
// GetBlogPostById - Retrieve a specific BlogPost

func (config *Config) GetBlogPostById(w http.ResponseWriter, r *http.Request) {

	blogPostId, err := strconv.ParseInt(r.URL.Query().Get("blogpostId"), 10, 64)
	if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	blogPost, err := config.dbSource.Blog.GetBlogPostById(blogPostId)
	if err == blog.ErrPostNotFound {
		log.Error(err)
		ErrorResponse(r, w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		log.Error(err)
		ErrorResponse(r, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, blogPost)
	render.Status(r, http.StatusOK)
}
