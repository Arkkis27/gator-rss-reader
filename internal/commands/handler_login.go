package commands

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/arkkis27/gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: login <username>")
	}
	username := cmd.Args[0]
	user, err := s.DB.GetUserByName(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("error: user '%s' does not exist\n", username)
			os.Exit(1)
		}
		return fmt.Errorf("error querying database: %w", err)
	}
	if err := s.Config.SetUser(user.Name); err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}
	fmt.Println("Logged in as:", user.Name)
	return nil
}
