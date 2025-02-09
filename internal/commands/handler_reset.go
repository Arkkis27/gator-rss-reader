package commands

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/arkkis27/gator/internal/state"
	"github.com/pressly/goose/v3"
)

// HandlerReset deletes the database
// func HandlerReset(s *state.State, cmd Command) error {
// 	err := s.DB.DBReset(s.Ctx)
// 	if err != nil {
// 		fmt.Println("Error resetting database:", err)
// 		return err
// 	}
// 	fmt.Println("Database reset successfully")
// 	return nil
// }

func HandlerReset(s *state.State, cmd Command) error {
	// Current reset code
	err := s.DB.DBReset(s.Ctx)
	if err != nil {
		fmt.Println("Error resetting database:", err)
		return err
	}

	// Get the path to your module root
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../../../") // Adjust the number of "../" based on where this code lives
	migrationPath := filepath.Join(rootPath, "gator", "internal", "database", "sql", "schema")

	goose.SetLogger(goose.NopLogger()) // This will silence goose's output
	err = goose.Reset(s.RawDB, migrationPath)
	if err != nil {
		return err
	}

	err = goose.Up(s.RawDB, migrationPath)
	if err != nil {
		fmt.Println("Error running migrations:", err)
		return err
	}

	fmt.Println("Database reset successfully")
	return nil
}
