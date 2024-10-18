package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// POST : create
// PUT : update
// DELETE : delete
// GET : read

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

// methods
//

func main() {
	endpoint := "localhost:3000"

	// initial load
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World\n")
		w.Write([]byte("there is no post at the moment"))
	})

	// create a post
	http.HandleFunc("POST /create", createBlogPostHandler)
	// get a blogpost
	http.HandleFunc("GET /blogs",getBlogPostByTitleHandler)

	if err != nil {
		fmt.Println("SErver failed")
		return
	}
}

func createBlogPostHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w,"Blog title already exists",http.StatusBadRequest)
	 }

	// append post to our memory
	blogPosts[blogPost.Title] = blogPost

	// prints out structs with field names
	fmt.Fprintf(w, "%+v", blogPost)
}

func listBlogPostHandler(w http.ResponseWriter, r *http.Request){

}

func getBlogPostTitleHandler(w http.ResponseWriter, r *http.Request){

}

func getBlogPostByTitleHandler(w http.ResponseWriter, r http.Request){
	title := r.URL.Query().Get("title")
	if title == ""{
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	blogPost, ok := blogPosts[title]
	if !ok {
		http.Error(w,"BlogPost not found", http.StatusNotFound)
		return

	}

	json.NewEncoder(w).Encode(blogPost)

}

func updateSpecificBlogPostTitleHandler(w http.ResponseWriter, r http.Request){
	title := r.URL.Query().Get("title")
	if title == ""{
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	blogPost, ok := blogPosts[title]
	if !ok {
		http.Error(w,"BlogPost not found", http.StatusNotFound)
		return

	}
	var update BlogPost
	// read request body

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}


	blogPosts[title] = update

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(blogPost)

}
// create a blog post api endpoint 
// 1.Create and save blog post on the map
// 2.
// 3.Edit a blog post
// 4. 
// 5. 

// get - query parameter :: /blogs?title=best