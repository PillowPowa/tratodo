package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"tratodo/internal/domain/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) GetById(id int) (*models.Todo, error) {
	const op = "repository.todo.GetById"

	stmt, err := r.db.Prepare(`SELECT id, title, completed FROM todos WHERE id = ?`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	t := new(models.Todo)
	err = stmt.QueryRow(id).Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return t, nil
}

func (r *TodoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	const op = "repository.todo.Create"

	stmt, err := r.db.Prepare(`INSERT INTO todos (title, completed) VALUES (?, ?)`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	t := new(models.Todo)
	err = stmt.QueryRow(todo.Title, todo.Completed).Scan(&t.ID, &t.Title)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return t, nil
}

func (r *TodoRepository) Delete(id int) error {
	const op = "repository.todo.Delete"

	stmt, err := r.db.Prepare(`DELETE FROM todos WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
