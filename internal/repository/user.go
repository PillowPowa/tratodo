package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"tratodo/internal/domain/models"

	"github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetById(id int64) (*models.User, error) {
	const op = "repository.user.GetById"

	stmt, err := r.db.Prepare(`SELECT id, username, pass_hash, email FROM users WHERE id = ?`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	u := models.User{}
	err = stmt.QueryRow(id).Scan(&u.ID, &u.Username, &u.PassHash, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &u, nil
}

func (r *UserRepository) GetUniqueUser(username string, email string) (*models.User, error) {
	const op = "repository.user.GetByInitials"

	stmt, err := r.db.Prepare(`SELECT id, username, pass_hash, email FROM users WHERE username = ? OR email = ?`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	u := models.User{}
	err = stmt.QueryRow(username, email).Scan(&u.ID, &u.Username, &u.PassHash, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &u, nil
}

func (r *UserRepository) Create(user *models.User) (int64, error) {
	const op = "repository.user.Create"

	stmt, err := r.db.Prepare(`INSERT INTO users (username, pass_hash, email) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(user.Username, user.PassHash, user.Email)
	if err != nil {
		if errors.Is(err, sqlite3.ErrConstraintUnique) {
			return 0, ErrAlreadyExists
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *UserRepository) Delete(id int64) error {
	const op = "repository.user.Delete"

	stmt, err := r.db.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
