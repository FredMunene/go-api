package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type BlogPost struct {
	Title   string
	Content string
	Tags    []string
	Scope   Scope
}

type Scope struct {
	IsPublic   bool
	Author     string
	ViewsCount int
}

// map origin :: memory database
// map with value as struct type BlogPost
var blogPosts = make(map[string]BlogPost)


func CreateBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	var blogPost BlogPost

	// read content from request body into a new decoder
	decoder := json.NewDecoder(r.Body)
	// decode content into our blog struct
	err := decoder.Decode(&blogPost)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusInternalServerError)
	}
	// check for uniwue blog title

	_, ok := blogPosts[blogPost.Title]
	if ok {
		http.Error(w, "Blog title already exists", http.StatusBadRequest)
		return
	}

	// append post to our memory
	blogPosts[blogPost.Title] = blogPost

	// prints out structs with field names
	fmt.Fprintf(w, "%+v", blogPost)
}

func ListBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	titles := []string{}

	for _, blog := range blogPosts {
		titles = append(titles, blog.Title)
	}

	if len(titles) == 0 {
		http.Error(w, "no blog posts found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(titles)

	// fmt.Fprintf(w,"%v",titles)
}

func IncreaseCountHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve title of blog
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// check title is present
	blogPost, ok := blogPosts[title]
	if !ok {
		http.Error(w, "BlogPost not found", http.StatusNotFound)
		return

	}

	blogPost.Scope.ViewsCount++

	blogPosts[title] = blogPost

	// into writer, encode the contents
	json.NewEncoder(w).Encode(blogPost)
}

func GetBlogPostByTitleHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve title of blog
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// check title is present
	blogPost, ok := blogPosts[title]
	if !ok {
		http.Error(w, "BlogPost not found", http.StatusNotFound)
		return

	}

	// into writer, encode the contents
	if err := json.NewEncoder(w).Encode(blogPost); err != nil {
		http.Error(w, "BlogPost not found", http.StatusInternalServerError)
		return
	}
}

func UpdateSpecificBlogPostTitleHandler(w http.ResponseWriter, r *http.Request) {
	// get title
	// check title is provided
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// check if such post exists
	blogPost, ok := blogPosts[title]
	if !ok {
		http.Error(w, "BlogPost not found", http.StatusNotFound)
		return

	}
	var update BlogPost

	// read request body
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// update post
	blogPosts[title] = update
	blogPost = update

	// return ok status
	w.WriteHeader(http.StatusOK)

	// return updated content
	json.NewEncoder(w).Encode(blogPost)
}

func DeleteSpecificBlog(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	
	// retrieve blog
	_, ok := blogPosts[title]
	if !ok {
		http.Error(w, "No blog with such title", http.StatusNotFound)
	}

	delete(blogPosts, title)

	w.WriteHeader(http.StatusOK)
}

// // create a blog post api endpoint
// // 1. Create and save blog post on the map.
// // 2. List all blog titles.
// // 3. Edit a blog post by title.
// // 4. Increase viewcount.
// // 5. Delete a blog post.
