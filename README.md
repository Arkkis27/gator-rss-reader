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
   git clone https://github.com/arkkis27/gator-rss-reader
   cd gator-rss-reader/internal/database/sql/schema
   goose postgres "<your_db_url_as_in_config>" up
   go install
  ```

### Database setup
1. Create a new PostgreSQL database and user with all privileges:
   ```sql
   CREATE DATABASE gator_db;
   CREATE USER gator_user WITH PASSWORD 'your_password';
   GRANT USAGE, CREATE ON SCHEMA public TO gator_user;
   GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO gator_user;
   GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO gator_user;
   ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO gator_user;
    ```
2. Set database password as environment variable (for example to ~/.bashrc):
   ```bash
   export GATOR_DB_PASSWORD='your_password'
   ```
3. Run migrations using goose:
   ```bash
   # Navigate to the migrations directory
   cd gator/internal/database/sql/schema
   
   # Run migrations (change your info)
   goose postgres "postgres://gator_user:${GATOR_DB_PASSWORD}@<your_db_server_ip>:5432/gator_db?sslmode=disable" up
   ```

###
* After installing the Gator-RSS-Reader, run "gator" in the terminal.
* This will error out, but at the same time it creates default configuration file. 
* This is saved to your home folder (path ~/.gatorconfig.json). Open the file and you will see following:
```bash
{
  "db_url": "postgres://gator_user:${GATOR_DB_PASSWORD}@<your_db_server_ip>:5432/gator_db?sslmode=disable",
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

# List users
gator users
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


# License
This project is licensed under the MIT License.