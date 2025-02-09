package commands

import (
	"fmt"
	"time"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *state.State, cmd Command, user gen.User) error {
	// Check that the function call gets both arguments
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <feed_name> <feed_url>")
	}

	// Call the RSS to get correctly edited HTML
	_, err := s.Client.FetchFeed(s.Ctx, cmd.Args[1])
	if err != nil {
		return err
	}

	// Prepare new feed parameters
	newFeed := gen.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	// Call to create a new row to feeds table
	tableFeed, err := s.DB.CreateFeed(s.Ctx, newFeed)
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	} else {
		// Print all fields if success
		fmt.Printf("%+v\n", tableFeed)

	}
	newFollow := gen.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    tableFeed.ID,
	}
	_, err = s.DB.CreateFeedFollow(s.Ctx, newFollow)
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}
	return nil
}
