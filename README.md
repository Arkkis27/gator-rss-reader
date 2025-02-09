# Gator RSS Reader
A command-line RSS feed aggregator written in Go. Allows users to manage and read RSS feeds from multiple sources.


## Features
- User authentication
- RSS feed management
- Feed aggregation
- Command-line interface



## Installation
### Prerequisites

- Go 1.20 or higher
- PostgreSQL 14+
- sqlc (install with: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)
- goose (install with: `go install github.com/pressly/goose/v3/cmd/goose@latest`)

### Database Setup
1. Create a new PostgreSQL database named 'gator'
2. Run migrations using goose:
   ```bash
   goose postgres "your-database-url" up

### Install from source with git
```bash
git clone https://github.com/arkkis27/gator
cd gator
go install
```

###
After installing the Gator-RSS-Reader, run "gator" in the terminal. This will error out, but at the same time it creates default configuration file. This is saved to your home folder (path ~/.gatorconfig.json). Open the file and you will see following:
```bash
{
  "db_url": "postgres://user:pass@localhost:5432/gator?sslmode=disable",
  "current_user_name": "",
  "user_agent": "Gator-RSS-Reader/1.0"
}
```
Update:
 * db_url - your database url
 * user_agent - add your contact information. When making htttp request, the program will use this as User-Agent header.


## Usage

### User management
```bash
# Register a new user
gator register <username>

# Login
gator login <username>
```

### Feed management
```bash
# Add a new feed
gator addfeed "Feed Name" "https://example.com/feed.xml"

# List feeds
gator feeds

# Follow already existing feed
gator follow <url>

# Unfollow feed
gator unfollow <url>
```


## Development
### Project structure
```bash
.
├── cmd/            # Application entrypoints
├── internal/       # Private packages
│   ├── commands/   # CLI commands
│   ├── config/     # App config
│   ├── database/   # Database operations
│   ├── rss/        # RSS parsing
│   └── state/      # RSS parsing 
```




# License
This project is licensed under the MIT License - see the LICENSE file for details.