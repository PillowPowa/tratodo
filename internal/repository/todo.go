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

func (r *TodoRepository) GetAll(query *models.TodoQuery, userId int64) ([]models.Todo, error) {
	const op = "repository.todo.GetAll"

	orderBy := fmt.Sprintf("ORDER BY %s", withSort(query.SortBy))
	filter := withFilter(query)
	stmt, err := r.db.Prepare(
		fmt.Sprintf(`SELECT id, title, completed, user_id FROM todos WHERE user_id = ? %s %s`, filter, orderBy),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	todos := []models.Todo{}
	for rows.Next() {
		t := models.Todo{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.UserId); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		todos = append(todos, t)
	}

	return todos, nil
}

func withFilter(query *models.TodoQuery) string {
	switch query.Filter {
	case "completed":
		return "AND completed = 1"
	case "uncompleted":
		return "AND completed = 0"
	default:
		return ""
	}
}

func withSort(sortBy string) string {
	switch sortBy {
	case "oldest":
		return "created_at ASC"
	default:
		return "created_at DESC"
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
