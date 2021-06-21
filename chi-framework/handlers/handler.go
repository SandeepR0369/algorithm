package handlers

import (
	"chi-framework/models"
	"encoding/json"
	"fmt"
	"net/http"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article models.Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(&article)
}
