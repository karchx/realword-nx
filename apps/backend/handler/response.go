package handler

import (
	"time"

	"github.com/karchx/realword-nx/model"
	"github.com/karchx/realword-nx/user"
	"github.com/karchx/realword-nx/utils"
	uuid "github.com/satori/go.uuid"
)

type userResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
		Token    string  `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Bio = u.Bio
	r.User.Image = u.Image
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

type profileResponse struct {
	Profile struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
		Following bool    `json:"following"`
	} `json:"profile"`
}

func newProfileResponse(us user.Store, UserID uuid.UUID, u *model.User) *profileResponse {
	r := new(profileResponse)
	r.Profile.Username = u.Username
	r.Profile.Bio = u.Bio
	r.Profile.Image = u.Image
	r.Profile.Following, _ = us.IsFollower(u.ID, UserID)
	return r
}

type articleResponse struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Body        string    `json:"body" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	TagList     []string  `json:"tagList"`
	Author      struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
		Following bool    `json:"following"`
	} `json:"author"`
}

type singleArticleResponse struct {
	Article *articleResponse `json:"article"`
}

func newArticleResponse(userID uuid.UUID, a *model.Article) *singleArticleResponse {
	ar := new(articleResponse)
	ar.Slug = a.Slug
	ar.Title = a.Title
	ar.Description = a.Description
	ar.Body = a.Body
	ar.CreatedAt = a.CreatedAt
	ar.UpdatedAt = a.UpdatedAt
	ar.Author.Username = a.Author.Username
	ar.Author.Image = a.Author.Image
	ar.Author.Bio = a.Author.Bio
	ar.Author.Following = a.Author.FollowedBy(userID)
	for _, t := range a.Tags {
		ar.TagList = append(ar.TagList, t.Tag)
	}

	return &singleArticleResponse{ar}
}
