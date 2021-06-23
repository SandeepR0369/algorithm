package models

import (
	"strings"
)

type Article struct {
	ID     int    `json:"Serial ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

var Articles []Article

func ListArticles() []Article {
	Articles = []Article{
		{ID: 1, Title: "GoLang", Author: "Alan"},
		{ID: 2, Title: "Python", Author: "William"},
		{ID: 3, Title: "Scala", Author: "Martin"},
		{ID: 4, Title: "Shell Script", Author: "Jason"},
		{ID: 5, Title: "Java Script", Author: "Brian"},
	}

	Articles = append(Articles, Article{ID: 6, Title: "Power of Subconscious Mind", Author: "Joseph Murphy"})
	Articles = append(Articles, Article{ID: 7, Title: "Jaffa", Author: "Brahmi"})

	return Articles
}

func Filter(vs []Article, f func(Article) bool) Article {
	for _, v := range vs {
		if f(v) {
			return v
		}
	}
	return Article{}
}

func ArticleById(id int) Article {
	return Filter(Articles, func(v Article) bool {
		return v.ID == id
	})
}

func ArticleByTitle(title string) Article {
	return Filter(Articles, func(v Article) bool {
		return strings.ToLower(v.Title) == strings.ToLower(title)
	})
}

func ArticleByAuthor(author string) Article {
	return Filter(Articles, func(v Article) bool {
		return strings.ToLower(v.Author) == strings.ToLower(author)
	})
}
