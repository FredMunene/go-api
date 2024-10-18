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

func main(){
	endpoint := "localhost:3000"

	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {	
		fmt.Fprintf(w, "Hello, World")
	})
	http.HandleFunc("POST /blogs", createBlogPostHandler)

	err := http.ListenAndServe(endpoint,nil)

	if err != nil {
		fmt.Println("SErver failed")
		return
	}
}

func createBlogPostHandler(w http.ResponseWriter, r *http.Request) {	
	var blogPost BlogPost

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&blogPost)
	if err != nil {
		http.Error(w,"Failed to decode request body", http.StatusInternalServerError)
	}

	blogPosts[blogPost.Title] = blogPost

	fmt.Fprintf(w,"%+v",blogPost)
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