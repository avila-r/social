package posts_test

import (
	"testing"

	p "github.com/avila-r/social/domain/posts"
	u "github.com/avila-r/social/domain/users"
)

var (
	users = []*u.User{
		{
			DisplayName: "Alice Test",
			Email:       "alice.test@example.com",
			Password:    "password123",
		},
	}
)

func Test_CreatePost(t *testing.T) {
	service := p.DefaultService

	service.Db.AutoMigrate(&p.Post{}, &u.User{})

	post := &p.Post{
		Sender:  users[0],
		Content: "This is a post!",
	}

	t.Cleanup(func() {
		service.DeleteByID(post.ID)

		// Before deleting any user, first asserts
		// that there's no posts related to them
		service.UserService.DeleteByID(post.SenderID)
	})

	if err := service.UserService.CreateUser(users[0]); err != nil {
		t.Errorf("error while trying to create user - %v", err)
	}

	if err := service.CreatePost(post); err != nil {
		t.Errorf("error while trying to create post - %v", err)
	}
}
