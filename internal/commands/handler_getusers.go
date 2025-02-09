package commands

import (
	"context"
	"fmt"

	"github.com/arkkis27/gator/internal/state"
)

// HandlerGetUsers gets all the users and marks the current login
func HandlerGetUsers(s *state.State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error getting users:", err)
		return err
	}
	for _, user := range users {
		if user == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}
