package entities

type Article struct {
	Id    string
	Title string
	Body  string
}

func NewArticle(id string, title string, body string) Article {
	return Article{
		Id:    id,
		Title: title,
		Body:  body,
	}
}
