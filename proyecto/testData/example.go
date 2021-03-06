package testData

import (
	"time"

	"example.com/mittaus/ddd-example/domain"
)

var rickBio = "Rick biography string"
var janeImg = "jane img link"

func User(name string) domain.User {
	switch name {
	case "rick":
		return rick
	default:
		return jane
	}
}
func Article(name string) domain.Article {
	switch name {
	default:
		return janeArticle
	}
}

var rick = domain.User{
	Name:      "rick",
	Email:     "rick@what.com",
	Bio:       &rickBio,
	ImageLink: nil,
	Password:  "password",
}

var jane = domain.User{
	Name:      "jane",
	Email:     "jane@what.com",
	Bio:       nil,
	ImageLink: &janeImg,
	Password:  "password",
}

const TokenPrefix = "Token "

var janeArticle = domain.Article{
	Slug:        "how-to-train-your-dragon",
	Title:       "articleTitle",
	Description: "description",
	Body:        "body",
	TagList:     []string{"tagList"},
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
	FavoritedBy: []domain.User{rick},
	Author:      jane,
	Comments: []domain.Comment{
		{ID: 123,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Body:      "commentBody",
			Author:    rick,
		},
	},
}
