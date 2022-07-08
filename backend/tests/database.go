package tests

import (
	"database/sql"
	"fmt"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/golang-migrate/migrate/v4"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func SetupDB(ctx *DockerContext) (*sql.DB, string) {
	// pulls an image, creates a container based on it and runs it
	resource, err := ctx.pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
		Networks: []*dockertest.Network{ctx.network},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.L.Fatal().Msgf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://postgres:secret@%s/dbname?sslmode=disable", hostAndPort)
	dockerIp := resource.GetIPInNetwork(ctx.network)
	dockerUrl := fmt.Sprintf("host=%s user=postgres dbname=dbname password=secret sslmode=disable", dockerIp)

	//log.L.Debug().Msgf("Connecting to database on url: %s", databaseUrl)

	err = resource.Expire(120) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		log.L.Error().Err(err).Msg("Failed to set expiration")
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	ctx.pool.MaxWait = 120 * time.Second
	var db *sql.DB
	if err = ctx.pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.L.Fatal().Msgf("Could not connect to docker: %s", err)
	}

	m, err := migrate.New(
		"file://../migrations",
		databaseUrl)
	if err != nil {
		log.L.Fatal().Err(err).Msg("Failed")
	}

	err = m.Up()
	if err != nil {
		log.L.Fatal().Err(err).Msg("Failed")
	}

	dbx := sqlx.NewDb(db, "postgres")
	_, err = dbx.Exec("INSERT INTO public.languages (code, name) VALUES ('no', 'Norwegian')")
	if err != nil {
		log.L.Fatal().Err(err).Msg("Failed")
	}

	return db, dockerUrl
}
