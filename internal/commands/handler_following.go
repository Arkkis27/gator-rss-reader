package commands

import (
	"fmt"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
)

func HandlerFollowing(s *state.State, cmd Command, user gen.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following command takes no arguments")
	}

	// Get all feednames the current user is following
	follows, err := s.DB.GetFeedFollowsForUser(s.Ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to get follows: %w", err)
	}
	for _, followrow := range follows {
		// Print feednames
		fmt.Printf("%+v\n", followrow.Feedname)
	}
	return nil
}
