package main

import (
	"chi-framework/handlers"
	"chi-framework/helpers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.Welcome)
	r.Get("/articles", handlers.ArticleList)
	r.Get("/special", handlers.Special)
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(helpers.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/post", handlers.CreatePost)
	})
	r.Get("/articles/id/{id}", handlers.ArticleId)
	r.Get("/articles/title/{title}", handlers.ArticleTitle)
	r.Get("/articles/author/{author}", handlers.ArticleAuthor)

	log.Fatal(http.ListenAndServe(":8081", r))
}
