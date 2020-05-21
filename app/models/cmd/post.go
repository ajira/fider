package cmd

import (
	"time"

	"github.com/getfider/fider/app/models"
	"github.com/getfider/fider/app/models/enum"
)

type AddNewPost struct {
	Title       string
	Description string

	Result *models.Post
}

type UpdatePost struct {
	Post                       *models.Post
	Title                      string
	Description                string
	IsPublic                   bool
	EstimatedDateForCompletion time.Time

	Result *models.Post
}

type SetPostResponse struct {
	Post   *models.Post
	Text   string
	Status enum.PostStatus
}
