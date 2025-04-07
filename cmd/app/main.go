package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DropsWeb/todogo/internal/database/postgres"
	"github.com/DropsWeb/todogo/internal/services"
	"github.com/DropsWeb/todogo/pkg/swagger"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/pressly/goose/v3"

	"github.com/go-openapi/loads"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../../.env")
	viper.ReadInConfig()

	connectData := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		viper.Get("POSTGRES_USER"),
		viper.Get("POSTGRES_PASSWORD"),
		viper.Get("POSTGRES_HOST"),
		viper.Get("POSTGRES_PORT"),
		viper.Get("POSTGRES_DBNAME"),
	)

	db, err := goose.OpenDBWithDriver("pgx", connectData)
	if err != nil {
		log.Fatalln(err)

	}
	if err := goose.Up(db, viper.GetString("GOOSE_MIGRATION_DIR")); err != nil {
		log.Fatalln(err)
	}

	connectDataPool := connectData + fmt.Sprintf(" pool_max_conns=%s", viper.Get("POSTGRES_MAX_POOL"))

	dbpool, err := pgxpool.New(context.Background(), connectDataPool)
	if err != nil {
		log.Fatalln(err)
	}

	defer dbpool.Close()

	connectDB := postgres.PgxConnect{DB: dbpool}

	swagger.Impl = &services.Server{
		ConnDb: &connectDB,
	}

	swaggerSpec, err := loads.Embedded(swagger.SwaggerJSON, swagger.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	server := swagger.NewServer(ops.NewTodoAPIAPI(swaggerSpec))

	defer server.Shutdown()

	server.ConfigureAPI()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:           300,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	s := http.Server{
		Addr:    ":" + viper.GetString("APP_PORT"),
		Handler: c.Handler(server.GetHandler()),
	}

	err = s.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}

}
