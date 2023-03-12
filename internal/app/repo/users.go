package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/evt/immulogapi/internal/app/models"
	"github.com/evt/immulogapi/internal/pkg/immudb"
)

const (
	testUserLogin    = "test"
	testUserPassword = "test"
)

type ImmuUserRepo struct {
	db *immudb.DB
}

// NewImmuUserRepo creates a new instance of ImmuDB repository
func NewImmuUserRepo(c *immudb.DB) *ImmuUserRepo {
	return &ImmuUserRepo{db: c}
}

func (r *ImmuUserRepo) CreateUsersTable(ctx context.Context) error {
	sql := `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER AUTO_INCREMENT,
    login VARCHAR[64],
    pass VARCHAR[256],
    PRIMARY KEY (id)
);`
	_, err := r.db.ExecContext(ctx, sql)
	if err != nil {
		return fmt.Errorf("db.ExecContext failed: %w", err)
	}
	return nil
}

func (r *ImmuUserRepo) CreateTestUser(ctx context.Context) error {
	testUser, err := r.GetUser(ctx, testUserLogin, testUserPassword)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		spew.Dump(err)
		log.Printf("failed to get user from repository: %v", err)
		return fmt.Errorf("GetUser failed: %w", err)
	}
	if testUser != nil {
		return nil
	}

	_, err = r.db.ExecContext(ctx, "INSERT INTO users(login, pass) VALUES (?, ?)", testUserLogin, testUserPassword)
	if err != nil {
		return fmt.Errorf("db.ExecContext failed: %w", err)
	}
	return nil
}

func (r *ImmuUserRepo) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	log.Printf("repo: create user: %v", user)

	res, err := r.db.ExecContext(ctx, "INSERT INTO users(login, pass) VALUES (?, ?)", user.Login, user.Pass)
	if err != nil {
		return 0, fmt.Errorf("db.ExecContext failed: %w", err)
	}
	userID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("db.LastInsertId failed: %w", err)
	}

	return userID, nil
}

func (r *ImmuUserRepo) GetUser(ctx context.Context, login, pass string) (*models.User, error) {
	log.Printf("repo: get user")

	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT id, login, pass FROM users").Scan(&user.ID, &user.Login, &user.Pass)
	if err != nil {
		return nil, fmt.Errorf("db.QueryRowContext failed: %w", err)
	}

	return &user, nil
}
