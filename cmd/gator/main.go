package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver

	"github.com/arkkis27/gator/internal/commands"
	"github.com/arkkis27/gator/internal/config"
	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/rss"
	"github.com/arkkis27/gator/internal/state"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading config: %s\n", err)
		os.Exit(1)
	}
	// Open database connection
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "error pinging database: %s\n", err)
		os.Exit(1)
	}

	// Create context
	ctx := context.Background()

	// Create state with database and config
	s := &state.State{
		Config: cfg,
		DB:     gen.New(db), // Attach
		RawDB:  db,
		Ctx:    ctx,
		Client: rss.NewClient(10*time.Second, cfg.UserAgent), //timeout lenght for http calls
	}

	// Initialize commands
	cmds := commands.Commands{
		Functions: make(map[string]func(*state.State, commands.Command) error),
	}

	// Register handlers
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerGetUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))
	cmds.Register("browse", commands.MiddlewareLoggedIn(commands.HandlerBrowse))

	// Validate args
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: not enough arguments")
		os.Exit(1)
	}

	// Create command from args
	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	// Run the program
	if err := cmds.Run(s, cmd); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
