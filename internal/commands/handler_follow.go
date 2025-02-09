package commands

import (
	"fmt"
	"time"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
	"github.com/google/uuid"
)

func HandlerFollow(s *state.State, cmd Command, user gen.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	feed, err := s.DB.GetFeedByURL(s.Ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error on getting feeds: %v", err)
	}

	// Prepare new feedfollow parameters
	newFollow := gen.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	// Call to create a new row to feedfollow table
	rowFF, err := s.DB.CreateFeedFollow(s.Ctx, newFollow)
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	} else {
		// Print all fields if success
		fmt.Printf("Feed: %s, CurrentUser: %s\n", rowFF.Feedname, rowFF.Username)
		return nil
	}
}
