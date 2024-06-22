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

func (r *TodoRepository) GetById(id int64) (*models.Todo, error) {
	const op = "repository.todo.GetById"

	stmt, err := r.db.Prepare(`SELECT id, title, completed, user_id FROM todos WHERE id = ?`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	t := models.Todo{}
	err = stmt.QueryRow(id).Scan(&t.ID, &t.Title, &t.Completed, &t.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &t, nil
}

func (r *TodoRepository) Create(todo *models.POSTTodo, userId int64) (int64, error) {
	const op = "repository.todo.Create"

	stmt, err := r.db.Prepare(`INSERT INTO todos (title, user_id) VALUES (?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(todo.Title, userId)
	if err != nil {
		// TEMP: bad solution, but it works :D
		if err.Error() == "FOREIGN KEY constraint failed" {
			return 0, ErrRef
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *TodoRepository) Delete(id int64) error {
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
