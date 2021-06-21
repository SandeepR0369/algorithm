package main

import (
	"chi-framework/handlers"
	"chi-framework/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
)

//var Articles []Article

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

	log.Fatal(http.ListenAndServe(":8081", r))
}
