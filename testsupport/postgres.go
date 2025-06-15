package testsupport

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type TestPostgres struct {
	DB         *sqlx.DB
	Container  *postgres.PostgresContainer
	ConnString string
}

var (
	once        sync.Once
	testPG      *TestPostgres
	initErr     error
	shutdownCtx context.Context
	shutdown    context.CancelFunc
)

// GetTestPostgres initializes and returns a singleton TestPostgres instance.
func GetTestPostgres() (*TestPostgres, error) {
	once.Do(func() {
		shutdownCtx, shutdown = context.WithCancel(context.Background())
		testPG, initErr = startPostgresContainer(shutdownCtx)
	})
	return testPG, initErr
}

// startPostgresContainer boots a PostgreSQL container and applies test_db_init.sql.
func startPostgresContainer(ctx context.Context) (*TestPostgres, error) {
	container, err := postgres.Run(
		ctx,
		"postgres:latest",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		postgres.WithInitScripts("../postgres/test_db_init.sql"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to get connection string: %w", err)
	}

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 30; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		container.Terminate(ctx)
		return nil, fmt.Errorf("database never became ready: %w", err)
	}

	return &TestPostgres{
		DB:         db,
		Container:  container,
		ConnString: connStr,
	}, nil
}

// StopTestPostgres can be deferred from TestMain or similar to ensure cleanup.
func StopTestPostgres() {
	if shutdown != nil {
		shutdown()
	}
	if testPG != nil && testPG.Container != nil {
		if err := testPG.Container.Terminate(context.Background()); err != nil {
			log.Printf("error terminating container: %v", err)
		}
	}
}
