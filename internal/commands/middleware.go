package commands

import (
	"fmt"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
)

func MiddlewareLoggedIn(handler func(s *state.State, cmd Command, user gen.User) error) func(*state.State, Command) error {
	return func(s *state.State, cmd Command) error {
		user, err := s.DB.GetUserByName(s.Ctx, s.Config.CurrentUserName)
		if err != nil {
			fmt.Println("Error getting user:", err)
			return err
		}
		return handler(s, cmd, user)
	}
}
