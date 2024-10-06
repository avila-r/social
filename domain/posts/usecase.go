package posts

// ListPosts retrieves all the posts from the database.
func (s *PostService) ListPosts() ([]Post, error) {
	var (
		posts []Post // Holds the result list of posts
	)

	// Query the database to find all posts
	// and store them in the 'posts' slice
	result := s.Db.Find(&posts)

	// Return the list of posts and any
	// error that occurred during the query
	return posts, result.Error
}

// CreatePost adds a new post to the database.
func (s *PostService) CreatePost(post *Post) error {
	if exists := s.UserService.ExistsById(post.SenderID); !exists {
		return ErrInvalidSenderID
	}

	// Insert the provided post into the database
	result := s.Db.Create(post)

	// Return any error that occurred during the insertion
	return result.Error
}

// FindByID retrieves a post from the database by its unique ID.
func (s *PostService) FindByID(id uint) (*Post, error) {
	var (
		post Post // Holds the result post
	)

	// Query the database for the first post with the given ID
	result := s.Db.Where("id = ?", id).First(&post)

	// Return the found post and any error
	// that occurred during the query
	return &post, result.Error
}

// DeleteByID deletes a post from the database by its unique ID.
func (s *PostService) DeleteByID(id uint) error {
	// Find the post by ID
	post, err := s.FindByID(id)

	if err != nil {
		// If an error occurred
		// (e.g., post not found),
		// return the error
		return err
	}

	// Delete the found post from the database
	result := s.Db.Delete(&post)

	// Return any error that occurred during the deletion
	return result.Error
}
