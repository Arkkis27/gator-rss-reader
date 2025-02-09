package commands

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/rss"
	"github.com/arkkis27/gator/internal/state"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func HandlerAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs> (e.g. '1s', '1m', '1h')")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %v", err)
	}

	// Minimum duration check
	minDuration := 5 * time.Second
	if timeBetweenReqs < minDuration {
		return fmt.Errorf("time between requests must be at least %v", minDuration)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		// Fetch RSS feed
		rssFeed, rssFeedID, err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feeds: %w", err)
		}
		// Save feed posts into the database
		err = savePostsToDB(s, rssFeed, rssFeedID)
		if err != nil {
			return fmt.Errorf("error saving posts to database: %w", err)
		}
	}
}

func scrapeFeeds(s *state.State) (*rss.RSSFeed, uuid.UUID, error) {
	// Get the next feed
	feed, err := s.DB.GetNextFeedToFetch(s.Ctx)
	if err != nil {
		return nil, uuid.UUID{}, err
	}

	// Mark the feed fetched
	err = s.DB.MarkFeedFetched(s.Ctx, feed.ID)
	if err != nil {
		return nil, uuid.UUID{}, err
	}

	// Fetch the feed content
	rssFeed, err := s.Client.FetchFeed(s.Ctx, feed.Url)
	if err != nil {
		return nil, uuid.UUID{}, err
	}
	return rssFeed, feed.ID, nil
}

// savePostsToDB processes the feed items and saves them as posts in the database
func savePostsToDB(s *state.State, rssFeed *rss.RSSFeed, rssFeedID uuid.UUID) error {
	for _, item := range rssFeed.Channel.Item {
		if item.Title == "" {
			fmt.Printf("Post with URL %s is missing a title, skipping.\n", item.Link)
			continue
		}
		pubTime, _ := parsePubDate(item.PubDate)
		postParams := gen.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: toNullString(item.Description),
			PublishedAt: pubTime, // Parse publication date properly
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			FeedID:      rssFeedID, // Ensure FeedID is part of rssFeed
		}

		_, err := s.DB.CreatePost(s.Ctx, postParams)
		if err != nil {
			// Ignore duplicate URL error, but log others
			if isDuplicateURLError(err) { // Example utility to detect duplicate errors
				fmt.Printf("Post with URL %s already exists, skipping\n", item.Link)
				continue
			}
			return fmt.Errorf("error saving post '%s': %w", item.Title, err)
		}
		fmt.Printf("Saved post: %s\n", item.Title)
	}
	return nil
}

// Converts a string to sql.NullString
func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// parsePubDate parses an RSS item's publication date string into a time.Time object.
func parsePubDate(date string) (time.Time, error) {
	// Attempt to parse using RFC1123 format (e.g., "Mon, 02 Jan 2006 15:04:05 MST")
	parsedTime, err := time.Parse(time.RFC1123, date)
	if err != nil {
		// If parsing fails, attempt RFC1123Z (e.g., includes numeric timezone offsets)
		parsedTime, err = time.Parse(time.RFC1123Z, date)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse pubDate: %s, error: %w", date, err)
		}
	}
	return parsedTime, nil
}

// IsDuplicateURLError checks if the error indicates a duplicate key violation on the "url" column.
func isDuplicateURLError(err error) bool {
	// Try to type-assert the error as a *pq.Error
	if pgErr, ok := err.(*pq.Error); ok {
		// Check if it's a unique violation (23505 error code)
		if pgErr.Code == "23505" {
			// Optional: Check if it's specifically the "posts_url_key" constraint
			// Adjust "posts_url_key" to match the name of your constraint
			if pgErr.Constraint == "posts_url_key" {
				return true
			}
		}
	}
	return false
}
