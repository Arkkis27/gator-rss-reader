package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/state"
	"github.com/google/uuid"
)

// HandlerRegister creates a new user
func HandlerRegister(s *state.State, cmd Command) error {
	// Ensure the command includes a username argument
	if len(cmd.Args) < 1 {
		return fmt.Errorf("no username provided")
	}
	username := cmd.Args[0]

	// Try to fetch the user by username
	_, err := s.DB.GetUserByName(context.Background(), username)

	// Handle the case where an unexpected error occurred
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error querying for user: %w", err)
	}

	if err == nil {
		// User already exists (no error, and user is valid)
		return fmt.Errorf("user '%s' already exists", username)
	}

	// Prepare new user parameters
	newUser := gen.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Attempt to create the new user in the database
	_, err = s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	// After successfully creating the user, update the config
	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}

	// Print confirmation of success
	fmt.Printf("User '%s' created successfully!\n", username)
	return nil
}
