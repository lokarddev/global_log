package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lokarddev/global_log/pkg/env"
	"io/ioutil"
	"log"
	"os"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	DbSchema string
	SSLMode  string
}

func InitDatabasePostgres() (*pgxpool.Pool, error) {
	cfg := PostgresConfig{
		Host:     env.DbHost,
		Port:     env.DbPort,
		Username: env.DbUser,
		Password: env.DbPass,
		DBName:   env.DbName,
		DbSchema: env.DbSchema,
		SSLMode:  env.DbSsl,
	}
	dsnDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.DbSchema, cfg.SSLMode)
	e, err := pgxpool.ParseConfig(dsnDB)
	if err != nil {
		return nil, err
	}
	e.MaxConns = int32(env.MaxCons)
	db, err := pgxpool.ConnectConfig(context.Background(), e)
	if err != nil {
		return nil, err
	}
	log.Printf("SUCCESSFUL CONNECTION TO DB[%s]\n", cfg.DBName)
	if env.AutoMigrate {
		if err = migrationsUp(); err != nil {
			log.Println(err.Error())
		}
	}
	err = createInitialData(db)
	if err != nil {
		log.Println(err)
	}
	return db, err
}

func migrationsUp() error {
	dsnMigrations := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.DbUser, env.DbPass, env.DbHost, env.DbPort, env.DbName, env.DbSsl)
	mDB, err := sql.Open("postgres", dsnMigrations)
	driver, err := postgres.WithInstance(mDB, &postgres.Config{
		MigrationsTable:       "schema_migrations",
		MigrationsTableQuoted: false,
		MultiStatementEnabled: false,
		DatabaseName:          env.DbName,
		SchemaName:            "public",
		StatementTimeout:      0,
		MultiStatementMaxSize: 0,
	})
	m, err := migrate.NewWithDatabaseInstance("file://migrations/", env.DbName, driver)
	err = m.Up()

	err = driver.Close()
	err = mDB.Close()
	return err
}

func createInitialData(db *pgxpool.Pool) error {
	path, err := os.Getwd()
	err = os.Chdir(fmt.Sprintf("%s/assets", path))
	if err != nil {
		return err
	}
	file, err := os.Open("initial_db.json")
	bytes, _ := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	resultMap := make(map[string]map[string]string)
	err = json.Unmarshal(bytes, &resultMap)
	for k, v := range resultMap {
		for _, j := range v {
			query := fmt.Sprintf("INSERT INTO %s (code, value) VALUES ($1, $1) ON CONFLICT (code) DO NOTHING", k)
			err = db.BeginFunc(context.Background(), func(tx pgx.Tx) error {
				_, err = db.Exec(context.Background(), query, j)
				return err
			})
		}
	}
	err = file.Close()
	return err
}
