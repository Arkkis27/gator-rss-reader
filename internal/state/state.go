package state

import (
	"context"
	"database/sql"

	"github.com/arkkis27/gator/internal/config"
	"github.com/arkkis27/gator/internal/database/sql/gen"
	"github.com/arkkis27/gator/internal/rss"
)

type State struct {
	Config *config.Config
	DB     *gen.Queries
	RawDB  *sql.DB
	Ctx    context.Context
	Client rss.Client
}
