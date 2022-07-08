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

	//dbx := sqlx.NewDb(db, "postgres")

	q := sqlc.New(db)
	episodes, err := q.GetEpisodes(ctx)
	assert.Nil(t, err)
	assert.NotEmpty(t, episodes)
}
