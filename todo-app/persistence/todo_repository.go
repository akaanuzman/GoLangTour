package persistence

import (
	"context"
	"time"
	"todoapp/domains"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type ITodoRepository interface {
	GetAllTodos() []domains.Todo
	GetTodoById(id int64) (domains.Todo, error)
	GetDoneOrUndoneTodos(isDone bool) []domains.Todo
	AddNewTodo(todo domains.Todo) error
	DeleteTodoById(id int64) error
	SignTodoAsDone(id int64, dueDate time.Time) error
	SingTodoAsUndone(id int64) error
}

type TodoRepository struct {
	dbPool *pgxpool.Pool
}

func NewTodoRepository(dbPool *pgxpool.Pool) ITodoRepository {
	return &TodoRepository{dbPool: dbPool}
}

func (todoRepository *TodoRepository) GetAllTodos() []domains.Todo {
	ctx := context.Background()
	todoRows, err := todoRepository.dbPool.Query(ctx, "SELECT * FROM todos")

	if err != nil {
		log.Error("Unable to get todos:", err)
		return []domains.Todo{}
	}

	return extractTodoRows(todoRows)
}

func (todoRepository *TodoRepository) GetTodoById(id int64) (domains.Todo, error) {
	ctx := context.Background()
	var todo domains.Todo
	err := todoRepository.dbPool.QueryRow(ctx, "SELECT * FROM todos WHERE id = $1", id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
		&todo.IsDone,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.DueDate,
	)

	if err != nil {
		log.Error("Unable to get todo by id:", err)
		return domains.Todo{}, err
	}

	return todo, nil
}

func (todoRepository *TodoRepository) GetDoneOrUndoneTodos(isDone bool) []domains.Todo {
	ctx := context.Background()
	todoRows, err := todoRepository.dbPool.Query(ctx, "SELECT * FROM todos WHERE is_done = $1", isDone)

	if err != nil {
		log.Error("Unable to get todos:", err)
		return []domains.Todo{}
	}

	return extractTodoRows(todoRows)
}

func (todoRepository *TodoRepository) AddNewTodo(todo domains.Todo) error {
	ctx := context.Background()
	_, err := todoRepository.dbPool.Exec(ctx,
		"INSERT INTO todos (title, description, is_done, created_at, updated_at, due_date) VALUES($1, $2, $3, $4, $5, $6)",
		todo.Title,
		todo.Description,
		todo.IsDone,
		todo.CreatedAt,
		todo.UpdatedAt,
		todo.DueDate,
	)
	if err != nil {
		log.Error("Unable to add new todo:", err)
		return err
	}
	return nil
}

func (todoRepository *TodoRepository) DeleteTodoById(id int64) error {
	ctx := context.Background()
	_, err := todoRepository.dbPool.Exec(ctx, "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		log.Error("Unable to delete todo by id:", err)
		return err
	}
	return nil
}

func (todoRepository *TodoRepository) SignTodoAsDone(id int64, dueDate time.Time) error {
	ctx := context.Background()
	_, err := todoRepository.dbPool.Exec(ctx, "UPDATE todos SET is_done = true, due_date = $1 WHERE id = $2", dueDate, id)
	if err != nil {
		log.Error("Unable to sign todo as done:", err)
		return err
	}
	return nil
}

func (todoRepository *TodoRepository) SingTodoAsUndone(id int64) error {
	ctx := context.Background()
	_, err := todoRepository.dbPool.Exec(ctx, "UPDATE todos SET is_done = false, due_date = NULL WHERE id = $1", id)
	if err != nil {
		log.Error("Unable to sign todo as undone:", err)
		return err
	}
	return nil
}

func extractTodoRows(todoRows pgx.Rows) []domains.Todo {
	todos := []domains.Todo{}
	var todo domains.Todo

	for todoRows.Next() {
		err := todoRows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.IsDone,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DueDate,
		)

		if err != nil {
			log.Error("Unable to scan todo:", err)
			return []domains.Todo{}
		}

		todos = append(todos, todo)
	}

	return todos
}
