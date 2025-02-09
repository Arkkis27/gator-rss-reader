package commands

import (
	"fmt"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
)

func HandlerUnfollow(s *state.State, cmd Command, user gen.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	feed, err := s.DB.GetFeedByURL(s.Ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error on getting feed: %v", err)
	}
	unFollowParams := gen.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.DB.Unfollow(s.Ctx, unFollowParams)
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}
	return nil
}
