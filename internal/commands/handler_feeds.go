package commands

import (
	"fmt"
	"strings"

	"github.com/arkkis27/gator/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {
	// Return if any args is used
	if len(cmd.Args) != 0 {
		fmt.Println("feeds doesn't take arguments")
		return nil
	}
	// Call to get the feeds table
	feedTable, err := s.DB.GetFeeds(s.Ctx)
	if err != nil {
		return fmt.Errorf("error on getting feeds: %v", err)
	}
	// Print column headers with formatting
	fmt.Printf("%-20s %-40s %-15s\n", "Feed Name:", "URL:", "User:")
	// Print a line below the headers
	fmt.Printf("%-20s %-40s %-15s\n", strings.Repeat("-", 20), strings.Repeat("-", 40), strings.Repeat("-", 15))

	for _, feedRow := range feedTable {
		user, err := s.DB.GetUserByID(s.Ctx, feedRow.UserID)
		if err != nil {
			fmt.Println("Error getting user:", err)
			return err
		}
		// Print rows with formatting
		fmt.Printf("%-20s %-40s %-15s\n", feedRow.Name, feedRow.Url, user.Name)
	}
	return nil
}
