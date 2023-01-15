package db

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/utils"
	"context"

	"github.com/gocql/gocql"
)

type IScylla interface {
	ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error
	ScanMapSlice(ctx context.Context, stmt string, arguments ...interface{}) ([]map[string]interface{}, error)
	Close()
}

type Scylla struct {
	session *gocql.Session
}

func NewScylla(c *models.DatabaseConfig) *Scylla {
	cluster := gocql.NewCluster(c.DatabaseHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.DatabaseUser,
		Password: c.DatabasePassword,
	}
	cluster.Keyspace = c.DatabaseKeyspace
	cluster.ConnectTimeout = cluster.ConnectTimeout * 5
	cluster.ProtoVersion = 4

	session, err := cluster.CreateSession()
	if err != nil {
		utils.Logger.Fatal("failed to create session", err)
		panic(session)
	}

	return &Scylla{
		session: session,
	}
}

func (s *Scylla) ScanMap(ctx context.Context, stmt string, results map[string]interface{}, arguments ...interface{}) error {
	q := s.session.Query(stmt, arguments...).WithContext(ctx)
	return q.MapScan(results)
}

func (s *Scylla) ScanMapSlice(ctx context.Context, stmt string, arguments ...interface{}) ([]map[string]interface{}, error) {
	q := s.session.Query(stmt, arguments...).WithContext(ctx)
	return q.Iter().SliceMap()
}

func (s *Scylla) Insert(ctx context.Context, stmt string, arguments ...interface{}) error {
	q := s.session.Query(stmt, arguments...).WithContext(ctx)
	return q.Exec()
}

func (s *Scylla) Close() {
	s.session.Close()
}
