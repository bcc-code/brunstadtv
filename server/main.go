package main

import (
	"database/sql"
	"fmt"

	db "admin.brunstad.tv/app/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {

	conn, err := sql.Open("pgx", "host=localhost user=postgres dbname=vod_v3 password=password sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	queries := db.New(conn)
	s := &Server{
		queries: queries,
		dbx:     sqlx.NewDb(conn, "pgx"),
	}

	r := gin.Default()
	r.GET("/medias", s.GetMedias)
	r.GET("/medias/:id", s.GetMedia)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
