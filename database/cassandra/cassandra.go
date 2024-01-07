package cassandra

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/migrate"
	"log"
)

func Connect() (*gocql.Session, error) {
	cluster := gocql.NewCluster("localhost")

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "dev",
		Password: "secret",
	}
	cluster.NumConns = 10
	cluster.Keyspace = "auth"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to cassandra")

	EnsureKeyspace(session, "auth")

	return session, nil
}

func CassandraMigrations(ctx context.Context, session *gocql.Session) {
	migrationsDir := "database/cassandra/migrations"

	if err := migrate.Migrate(ctx, session, migrationsDir); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
}

func EnsureKeyspace(session *gocql.Session, keyspace string) {
	query := `CREATE KEYSPACE IF NOT EXISTS ` + keyspace + ` WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};`
	if err := session.Query(query).Exec(); err != nil {
		log.Fatalf("Failed to create keyspace: %v", err)
	}
}
