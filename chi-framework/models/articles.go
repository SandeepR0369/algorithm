package models

type Article struct {
	ID    int    `json:"Serial ID"`
	Title string `json:"Title"`
	PNo   int    `json:"Page Number"`
}

var Articles []Article

func ListArticles() []Article {
	Articles = []Article{
		{ID: 1, Title: "GoLang", PNo: 4},
		{ID: 2, Title: "Python", PNo: 24},
		{ID: 3, Title: "Scala", PNo: 44},
		{ID: 4, Title: "Shell Script", PNo: 64},
		{ID: 5, Title: "Java Script", PNo: 74},
	}

	Articles = append(Articles, Article{ID: 6, Title: "Power of Subconscious Mind", PNo: 51})
	Articles = append(Articles, Article{ID: 7, Title: "Jaffa", PNo: 143})

	return Articles
}
