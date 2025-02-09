package commands

import (
	"fmt"
	"strconv"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
)

func HandlerBrowse(s *state.State, cmd Command, user gen.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: browse <limit (optional, default 2)>")
	}
	searchLimit := int32(2)
	if len(cmd.Args) == 1 {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("usage: browse <limit (optional, default 2)>")
		}
		searchLimit = int32(i)
	}

	getParams := gen.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  searchLimit,
	}
	post, err := s.DB.GetPostsForUser(s.Ctx, getParams)
	if err != nil {
		return fmt.Errorf("failed to fetch posts for user: %w", err)
	}
	if len(post) == 0 {
		fmt.Println("No posts available.")
		return nil
	}
	for _, p := range post {
		fmt.Printf("Title: %s\nURL: %s\nPublished At: %s\n\n", p.Title, p.Url, p.PublishedAt)
	}
	return nil
}
