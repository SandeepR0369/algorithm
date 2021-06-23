package handlers

import (
	"chi-framework/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var Articles []models.Article

func ArticleList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ListArticles())
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Chi-Framework POC")
}

func Special(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func ArticleId(w http.ResponseWriter, r *http.Request) {
	// get the id from request
	models.ListArticles()
	p := strings.Split(r.URL.Path, "/")
	id := ""
	if len(p) > 0 {
		id = p[3]
	} else {
		json.NewEncoder(w).Encode(models.Article{})
	}
	ID, _ := strconv.Atoi(id)
	a := models.ArticleById(ID)
	json.NewEncoder(w).Encode(a)

}

func ArticleTitle(w http.ResponseWriter, r *http.Request) {
	// get the id from request
	models.ListArticles()
	p := strings.Split(r.URL.Path, "/")
	title := ""
	if len(p) > 0 {
		title = p[3]
	} else {
		json.NewEncoder(w).Encode(models.Article{})
	}

	a := models.ArticleByTitle(title)
	json.NewEncoder(w).Encode(a)

}

func ArticleAuthor(w http.ResponseWriter, r *http.Request) {
	// get the id from request
	models.ListArticles()
	p := strings.Split(r.URL.Path, "/")
	author := ""
	fmt.Println(p, len(p))
	if len(p) > 0 {
		author = p[3]
	} else {
		json.NewEncoder(w).Encode(models.Article{})
	}

	a := models.ArticleByAuthor(author)
	json.NewEncoder(w).Encode(a)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article models.Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(&article)
}
