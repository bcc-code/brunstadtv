package tests

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/sqlc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerDatabase(t *testing.T) {
	ctx := context.Background()
	dockerContext := SetupDocker()
	db, _ := SetupDB(dockerContext)

	q := sqlc.New(db)
	episodes, _ := q.GetEpisodes(ctx)
	assert.NotEmpty(t, episodes)
}
