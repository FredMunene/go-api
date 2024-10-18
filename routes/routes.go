package routes

import (
	"go-api/handler"
	"go-api/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	blogPostRoutes(r)
	return r
}

func blogPostRoutes(r *mux.Router) {
	r.HandleFunc("/create", handler.CreateBlogPostHandler).Methods("POST")
	// get a blogpost
	r.HandleFunc("/blogs", handler.GetBlogPostByTitleHandler).Methods("GET")
	// list title of blogs
	r.HandleFunc("/list", handler.ListBlogPostHandler).Methods("GET")
	// list title of blogs
	r.HandleFunc("/count", handler.IncreaseCountHandler).Methods("GET")
	// edit a blog post
	r.HandleFunc("/edit", handler.UpdateSpecificBlogPostTitleHandler).Methods("PUT")
	// delete a blog post
	r.HandleFunc("/delete", middleware.RequireAuth(handler.DeleteSpecificBlog)).Methods("DELETE")
}
